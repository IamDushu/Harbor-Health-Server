package api

import (
	"net/http"

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
