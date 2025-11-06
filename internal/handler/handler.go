package handler

import (
	"log/slog"

	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	config "github.com/amenshenin/go_auth/internal/configs"
	middleware "github.com/amenshenin/go_auth/internal/handler/middleware/logger"
	"github.com/amenshenin/go_auth/internal/service"
)

type Handler struct {
	config  *config.Config
	logger  *slog.Logger
	service *service.Service
}

func NewHandler(cfg *config.Config, l *slog.Logger, s *service.Service) *Handler {
	return &Handler{
		config:  cfg,
		logger:  l,
		service: s,
	}
}

func (h *Handler) InitRouts() *gin.Engine {
	switch h.config.Enviremant {
	case config.EnvProd:
		gin.SetMode(gin.ReleaseMode)
	case config.EnvDev:
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(requestid.New())
	router.Use(middleware.Logger(h.logger))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}
	api := router.Group("/api", h.userIdentify)
	{
		api.GET("/test")
	}
	return router
}
