package api

import (
	"net/http"
	"time"

	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/IamDushu/Harbor-Health-Server/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createMemberRequest struct {
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	PhoneNumber    string    `json:"phone_number"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	AddressLineOne string    `json:"address_line_one"`
	AddressLineTwo string    `json:"address_line_two"`
	Insurance      string    `json:"insurance"`
	AcceptedTerms  bool      `json:"accepted_terms"`
}

type memberResponse struct {
	MemberID  uuid.UUID `json:"member_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func newMemberResponse(member db.Member) memberResponse {
	return memberResponse{
		MemberID:  member.MemberID,
		UserID:    member.UserID,
		CreatedAt: member.CreatedAt,
	}
}

func (s *Server) CreateMember(ctx *gin.Context) {
	var req createMemberRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	memArgs := db.CreateMemberArgs{
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		PhoneNumber:    req.PhoneNumber,
		Email:          authPayload.Email,
		DateOfBirth:    req.DateOfBirth,
		Gender:         req.Gender,
		AddressLineOne: req.AddressLineOne,
		AddressLineTwo: req.AddressLineTwo,
		Insurance:      req.Insurance,
		AcceptedTerms:  req.AcceptedTerms,
	}

	newMember, err := s.store.CreateMemberTx(ctx, memArgs, s.streamClient)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newMemberResponse(newMember)
	ctx.JSON(http.StatusOK, response)
}
