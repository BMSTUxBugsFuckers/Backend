package service

import (
	"candc/entity"
	"candc/internal/repository"
	"go.uber.org/zap"
)

type IServiceInstance interface {
	CreateInstance(instance entity.Instance) error
	GetInstance(id int64) ([]entity.Instance, error)
	GetAllInstances() ([]entity.Instance, error)
	UpdateInstance(instance entity.Instance) error
	DeleteInstance(id int64) error
}

type Service struct {
	IServiceInstance
}

func NewService(repo *repository.Repository, logger *zap.Logger) *Service {
	return &Service{
		IServiceInstance: NewInstanceService(repo, logger),
	}
}
