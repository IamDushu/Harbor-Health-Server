package api

import (
	"fmt"

	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/IamDushu/Harbor-Health-Server/internal/token"
	"github.com/IamDushu/Harbor-Health-Server/internal/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our api service.
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/api/registration/email", s.registerUser)
	router.POST("/api/registration/email/verify", s.verifyUser)
	router.POST("/api/tokens/renew_access", s.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))
	authRoutes.POST("/api/users", s.UpdateUser)
	authRoutes.POST("/api/members", s.CreateMember)

	s.router = router
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}
