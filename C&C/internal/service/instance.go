package service

import (
	"candc/entity"
	"candc/internal/repository"
	"go.uber.org/zap"
)

type InstanceService struct {
	repo   *repository.Repository
	logger *zap.Logger
}

func (i InstanceService) CreateInstance(instance entity.Instance) error {
	return i.repo.CreateInstance(instance)
}

func (i InstanceService) GetInstance(id int64) ([]entity.Instance, error) {
	return i.repo.GetInstance(id)
}

func (i InstanceService) GetAllInstances() ([]entity.Instance, error) {
	return i.repo.GetAllInstances()
}

func (i InstanceService) UpdateInstance(instance entity.Instance) error {
	return i.repo.UpdateInstance(instance)
}

func (i InstanceService) DeleteInstance(id int64) error {
	return i.repo.DeleteInstance(id)
}

func NewInstanceService(repo *repository.Repository, logger *zap.Logger) *InstanceService {
	return &InstanceService{
		repo:   repo,
		logger: logger,
	}
}
