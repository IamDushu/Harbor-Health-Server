package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetLocations(ctx *gin.Context) {

	locations, err := s.store.GetLocations(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, locations)
}
