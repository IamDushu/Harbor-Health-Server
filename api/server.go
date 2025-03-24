package api

import (
	"fmt"

	stream "github.com/GetStream/stream-chat-go/v5"
	db "github.com/IamDushu/Harbor-Health-Server/internal/db/sqlc"
	"github.com/IamDushu/Harbor-Health-Server/internal/token"
	"github.com/IamDushu/Harbor-Health-Server/internal/util"
	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our api service.
type Server struct {
	config       util.Config
	store        *db.Store
	tokenMaker   token.Maker
	router       *gin.Engine
	streamClient *stream.Client
}

func (s *Server) setupRouter() {
	router := gin.Default()

	router.POST("/api/registration/email", s.registerUser)
	router.POST("/api/registration/email/verify", s.verifyUser)
	router.POST("/api/tokens/renew_access", s.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))
	authRoutes.GET("/api/users", s.GetUser)
	authRoutes.GET("/api/locations", s.GetLocations)
	authRoutes.GET("/api/providers", s.GetProvidersFromLocation)
	authRoutes.GET("/api/visits", s.GetPendingVisits)
	authRoutes.GET("/api/visits/:visit_id", s.GetVisit)
	authRoutes.GET("/api/providers/:provider_id/availability", s.GetProviderAvailability)
	authRoutes.POST("/api/users", s.UpdateUser)
	authRoutes.POST("/api/members", s.CreateMember)
	authRoutes.POST("/api/visits", s.CreateVisit)

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

	client, err := stream.NewClient(config.StreamApiKey, config.StreamSecretKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create stream client: %w", err)
	}

	server := &Server{
		config:       config,
		store:        store,
		tokenMaker:   tokenMaker,
		streamClient: client,
	}

	server.setupRouter()
	return server, nil
}
