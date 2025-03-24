package api

import (
	"net/http"
	"time"

	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/IamDushu/Harbor-Health-Server/internal/token"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type createVisitRequest struct {
	UserID     uuid.UUID `json:"user_id"`
	ProviderID uuid.UUID `json:"provider_id"`
	LocationID uuid.UUID `json:"location_id"`
	Date       string    `json:"date"`
	StartTime  string    `json:"start_time"`
	Notes      string    `json:"notes"`
}

type visitResponse struct {
	VisitID     uuid.UUID `json:"visit_id"`
	ProviderID  uuid.UUID `json:"provider_id"`
	MembersID   uuid.UUID `json:"members_id"`
	LocationID  uuid.UUID `json:"location_id"`
	ScheduledAt time.Time `json:"scheduled_at"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

func newVisitResponse(visit db.Visit) visitResponse {
	return visitResponse{
		VisitID:     visit.VisitID,
		ProviderID:  visit.ProviderID,
		MembersID:   visit.MemberID,
		LocationID:  visit.LocationID,
		ScheduledAt: visit.ScheduledAt,
		Status:      visit.Status,
		CreatedAt:   visit.CreatedAt,
	}
}

func (s *Server) CreateVisit(ctx *gin.Context) {
	var req createVisitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	member, err := s.store.GetMember(ctx, req.UserID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Can't find member"})
		return
	}

	selectedDate, err := time.Parse("2006-01-02", req.Date) //YYYY-MM-DD
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid date format"})
		return
	}

	startTime, err := time.Parse("03:04 PM", req.StartTime)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid time format"})
		return
	}

	dayOfWeek := int32(selectedDate.Weekday())

	scheduledAt := selectedDate.Add(time.Hour*time.Duration(startTime.Hour()) + time.Minute*time.Duration(startTime.Minute()))

	availabilityExists, err := s.store.CheckProviderAvailability(ctx, db.CheckProviderAvailabilityParams{
		ProviderID: req.ProviderID,
		DayOfWeek:  dayOfWeek,
		StartTime:  startTime,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking provider availability"})
		return
	}

	if !availabilityExists {
		ctx.JSON(http.StatusConflict, gin.H{"error": "Selected time slot is unavailable"})
		return
	}

	visitArgs := db.CreateVisitArgs{
		ProviderID:  req.ProviderID,
		LocationID:  req.LocationID,
		MemberID:    member.MemberID,
		Notes:       req.Notes,
		ScheduledAt: scheduledAt,
	}

	newVisit, err := s.store.CreateVisitTx(ctx, visitArgs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := newVisitResponse(newVisit)
	ctx.JSON(http.StatusOK, response)
}

func (s *Server) GetPendingVisits(ctx *gin.Context) {

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := s.store.GetUser(ctx, authPayload.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	member, err := s.store.GetMember(ctx, user.UserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	visits, err := s.store.GetAllPendingVisitsWithProviderDetails(ctx, member.MemberID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, visits)
}

func (s *Server) GetVisit(ctx *gin.Context) {
	visitID := ctx.Param("visit_id")

	parsedVisitID, err := uuid.Parse(visitID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid visit ID format"})
		return
	}

	visitInfo, err := s.store.GetVisitInfo(ctx, parsedVisitID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, visitInfo)
}
