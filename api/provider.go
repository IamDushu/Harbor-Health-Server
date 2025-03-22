package api

import (
	"net/http"
	"time"

	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (s *Server) GetProvidersFromLocation(ctx *gin.Context) {

	locationID := ctx.DefaultQuery("location_id", "")
	if locationID == "" {
		ctx.JSON(400, gin.H{"error": "location_id is required"})
		return
	}

	locationUUID, err := uuid.Parse(locationID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid location_id format"})
		return
	}

	providers, err := s.store.GetProvidersFromLocation(ctx, locationUUID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, providers)
}

func (s *Server) GetProviderAvailability(ctx *gin.Context) {
	providerID := ctx.Param("provider_id")
	dateStr := ctx.DefaultQuery("date", "") // Expecting a date in YYYY-MM-DD format (example: "2025-04-01")

	if dateStr == "" {
		ctx.JSON(400, gin.H{"error": "date is required"})
		return
	}

	// Parse the provided date string into a Date object
	selectedDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "Invalid date format, use YYYY-MM-DD"})
		return
	}

	providerUUID, err := uuid.Parse(providerID)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "invalid provider_id format"})
		return
	}

	dayOfWeek := int32(selectedDate.Weekday())

	args := db.GetAvailableSlotsForProviderParams{
		ProviderID:  providerUUID,
		DayOfWeek:   dayOfWeek,
		ScheduledAt: selectedDate,
	}

	availableSlots, err := s.store.GetAvailableSlotsForProvider(ctx, args)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "could not retrieve availability"})
		return
	}

	// Format the times into a more readable format (e.g., "09:00 AM" or "09:00")
	var formattedSlots []map[string]interface{}
	for _, slot := range availableSlots {
		startTimeFormatted := slot.StartTime.Format("03:04 PM")
		endTimeFormatted := slot.EndTime.Format("03:04 PM")

		formattedSlot := map[string]interface{}{
			"day_of_week": slot.DayOfWeek,
			"start_time":  startTimeFormatted,
			"end_time":    endTimeFormatted,
		}
		formattedSlots = append(formattedSlots, formattedSlot)
	}

	ctx.JSON(200, gin.H{
		"provider_id":     providerID,
		"date":            selectedDate.Format("2006-01-02"),
		"available_slots": formattedSlots,
	})
}
