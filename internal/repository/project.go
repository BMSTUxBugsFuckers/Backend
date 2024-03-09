package repository

import (
	"backend/entity"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProjectDB struct {
	db *sqlx.DB
}

func (p ProjectDB) CreateProject() {
	//TODO implement me
	panic("implement me")
}

func (p ProjectDB) GetProject(id int64) (entity.Project, error) {
	var project entity.Project

	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", projectsTable)
	err := p.db.Get(&project, query, id)

	return project, err
}

func (p ProjectDB) GetAllProjects() {
	//TODO implement me
	panic("implement me")
}

func (p ProjectDB) UpdateProject() {
	//TODO implement me
	panic("implement me")
}

func (p ProjectDB) DeleteProject() {
	//TODO implement me
	panic("implement me")
}

func NewProjectDB(db *sqlx.DB) *ProjectDB {
	return &ProjectDB{db: db}
}
