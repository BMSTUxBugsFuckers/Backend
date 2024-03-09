package handler

import "backend/internal/usecase"

type Handler struct {
	services *usecase.UseCase
}

func NewHandler(services *usecase.UseCase) *Handler {
	return &Handler{services: services}
}
