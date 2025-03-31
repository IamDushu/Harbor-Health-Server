package api

import "github.com/gin-gonic/gin"

func (s *Server) healthCheck(ctx *gin.Context) {
	err := s.store.HealthCheck(ctx) // optional DB check
	if err != nil {
		ctx.JSON(500, gin.H{"status": "error", "db": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"status": "ok"})
}
