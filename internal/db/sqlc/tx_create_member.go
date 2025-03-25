package db

import (
	"context"
	"fmt"
	"time"

	stream_chat "github.com/GetStream/stream-chat-go/v5"
	"github.com/google/uuid"
)

type CreateMemberArgs struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	PhoneNumber    string    `json:"phone_number"`
	Email          string    `json:"email"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	AddressLineOne string    `json:"address_line_one"`
	AddressLineTwo string    `json:"address_line_two"`
	Insurance      string    `json:"insurance"`
	AcceptedTerms  bool      `json:"accepted_terms"`
}

func (s *Store) CreateMemberTx(ctx context.Context, memberArgs CreateMemberArgs, streamClient *stream_chat.Client) (Member, error) {

	var newMember Member

	err := s.execTx(ctx, func(q *Queries) error {
		userArgs := UpdateUserParams{
			FirstName:   memberArgs.FirstName,
			LastName:    memberArgs.LastName,
			PhoneNumber: memberArgs.PhoneNumber,
			IsOnboarded: true,
			Email:       memberArgs.Email,
		}

		user, err := q.UpdateUser(ctx, userArgs)
		if err != nil {
			return err
		}

		memArgs := CreateMemberParams{
			MemberID:       uuid.New(),
			UserID:         user.UserID,
			Gender:         memberArgs.Gender,
			DateOfBirth:    memberArgs.DateOfBirth,
			Insurance:      memberArgs.Insurance,
			AddressLineOne: memberArgs.AddressLineOne,
			AddressLineTwo: memberArgs.AddressLineTwo,
			AcceptedTerms:  memberArgs.AcceptedTerms,
		}

		member, err := q.CreateMember(ctx, memArgs)
		if err != nil {
			return err
		}

		_, err = streamClient.UpsertUsers(ctx, &stream_chat.User{
			ID:    user.UserID.String(),
			Name:  user.FirstName,
			Image: user.ImageUrl.String,
		})
		if err != nil {
			return fmt.Errorf("failed to upsert user in Stream: %w", err)
		}

		newMember = member
		return nil
	})

	return newMember, err
}
