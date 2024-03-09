package repository

import (
	"backend/entity"
	"github.com/jmoiron/sqlx"
)

type Project interface {
	CreateProject()
	GetProject(id int64) (entity.Project, error)
	GetAllProjects()
	UpdateProject()
	DeleteProject()
}

type Repository struct {
	Project
}

func NewRepo(db *sqlx.DB) *Repository {
	return &Repository{
		Project: NewProjectDB(db),
	}
}
