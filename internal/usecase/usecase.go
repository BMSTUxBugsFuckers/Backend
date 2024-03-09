package usecase

import (
	"backend/entity"
	"backend/internal/repository"
)

type Project interface {
	CreateProject()
	GetProject(id int64) (entity.Project, error)
	GetAllProjects()
	UpdateProject()
	DeleteProject()
}

type UseCase struct {
	Project
}

func NewUseCase(repo *repository.Repository) *UseCase {
	return &UseCase{
		Project: NewProjectUseCase(repo),
	}
}
