package repository

import (
	"candc/entity"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type IRepoInstance interface {
	CreateInstance(instance entity.Instance) error
	GetInstance(id int64) ([]entity.Instance, error)
	GetAllInstances() ([]entity.Instance, error)
	UpdateInstance(instance entity.Instance) error
	DeleteInstance(id int64) error
}

type Repository struct {
	IRepoInstance
}

func NewRepo(db *sqlx.DB, logger *zap.Logger) *Repository {
	return &Repository{
		IRepoInstance: NewInstanceDB(db, logger),
	}
}
