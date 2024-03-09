package usecase

import (
	"backend/entity"
	"backend/internal/repository"
)

type ProjectUseCase struct {
	repo *repository.Repository
}

func (p ProjectUseCase) CreateProject() {
	//TODO implement me
	panic("implement me")
}

func (p ProjectUseCase) GetProject(id int64) (entity.Project, error) {
	return p.repo.Project.GetProject(id)
}

func (p ProjectUseCase) GetAllProjects() {
	//TODO implement me
	panic("implement me")
}

func (p ProjectUseCase) UpdateProject() {
	//TODO implement me
	panic("implement me")
}

func (p ProjectUseCase) DeleteProject() {
	//TODO implement me
	panic("implement me")
}

func NewProjectUseCase(repo *repository.Repository) *ProjectUseCase {
	return &ProjectUseCase{repo: repo}
}
