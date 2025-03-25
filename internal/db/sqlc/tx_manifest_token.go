package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/IamDushu/Harbor-Health-Server/internal/util"
	"github.com/google/uuid"
)

func (s *Store) ManifestTokenTx(ctx context.Context, verificationRecord EmailVerification) error {
	err := s.execTx(ctx, func(q *Queries) error {
		_, err := q.UpdateVerifyRecordInvalid(ctx, verificationRecord.VerificationID)
		if err != nil {
			return err
		}

		if verificationRecord.Purpose == util.SIGNUP {
			user_image := fmt.Sprintf("https://api.dicebear.com/9.x/open-peeps/png?size=96&radius=50&backgroundColor=63daae&clothingColor=000000&accessoriesProbability=0&face=calm,cute,smile&skinColor=f1f0f5&translateX=-10&seed=123%s", verificationRecord.VerificationID)
			_, err = q.CreateUser(ctx, CreateUserParams{
				UserID:   uuid.New(),
				Email:    verificationRecord.Email,
				ImageUrl: sql.NullString{String: user_image, Valid: true},
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}
