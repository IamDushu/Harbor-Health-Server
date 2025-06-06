package api

import (
	"database/sql"
	"net/http"
	"time"

	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/IamDushu/Harbor-Health-Server/internal/token"
	"github.com/IamDushu/Harbor-Health-Server/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createUserRequest struct {
	Email       string `json:"email" binding:"required,email"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

type userResponse struct {
	UserID      uuid.UUID      `json:"user_id"`
	Email       string         `json:"email"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	ImageUrl    sql.NullString `json:"image_url"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	IsOnboarded bool           `json:"is_onboarded"`
	CreatedAt   time.Time      `json:"created_at"`
}

type updateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		UserID:      user.UserID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		ImageUrl:    user.ImageUrl,
		PhoneNumber: user.PhoneNumber,
		IsOnboarded: user.IsOnboarded,
		CreatedAt:   user.CreatedAt,
	}
}

func (s *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	req.Email = util.NormalizeEmail(req.Email)

	arg := db.CreateUserParams{
		UserID:      uuid.New(),
		Email:       req.Email,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
	}

	user, err := s.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

func (s *Server) UpdateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	// phoneDetails, err := util.VerifyPhone(req.PhoneNumber, s.config.TwillioAccountSID, s.config.TwillioAuthToken)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	// 	return
	// }
	// if !phoneDetails.Valid {
	// 	ctx.JSON(http.StatusBadRequest, errorResponse(fmt.Errorf("invalid phone number: %v", phoneDetails.ValidationErrors)))
	// 	return
	// }

	arg := db.UpdateUserParams{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		PhoneNumber: req.PhoneNumber,
		Email:       authPayload.Email,
	}

	user, err := s.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}

func (s *Server) GetUser(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := s.store.GetUser(ctx, authPayload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newUserResponse(user)
	ctx.JSON(http.StatusOK, response)
}
