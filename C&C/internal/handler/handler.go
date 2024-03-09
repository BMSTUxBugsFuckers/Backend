package handler

import (
	"candc/internal/usecase"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *usecase.UseCase
}

func NewHandler(services *usecase.UseCase) *Handler {
	return &Handler{services: services}
}

func InitRoute() *fiber.App {
	router := fiber.New()
	return router
}
