package handler

import (
	"candc/internal/service"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

type Handler struct {
	services *service.Service
	logger   *zap.Logger
}

func NewHandler(services *service.Service, logger *zap.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func InitRoute(handlers *Handler) *fiber.App {
	router := fiber.New()
	api := router.Group("/api")
	api.Get("/instance/:id", handlers.GetInstance)
	return router
}
