package usecase

import (
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/entity"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/repository"
)

type SensorUsecase interface {
	CreateSensorData(data *entity.SensorData) error
	FindByIDs(id1 string, id2 int) (*entity.SensorData, error)
	FindByTimeRange(start, end int64) ([]entity.SensorData, error)
	DeleteByIDs(id1 string, id2 int) error
	UpdateSensorData(data *entity.SensorData) error
}

type sensorUsecase struct {
	repo repository.SensorRepository
}

func NewSensorUsecase(repo repository.SensorRepository) *sensorUsecase {
	return &sensorUsecase{repo: repo}
}

func (uc *sensorUsecase) CreateSensorData(data *entity.SensorData) error {
	return uc.repo.Create(data)
}

func (uc *sensorUsecase) FindByIDs(id1 string, id2 int) (*entity.SensorData, error) {
	return uc.repo.FindByIDs(id1, id2)
}

func (uc *sensorUsecase) FindByTimeRange(start, end int64) ([]entity.SensorData, error) {
	return uc.repo.FindByTimeRange(start, end)
}

func (uc *sensorUsecase) DeleteByIDs(id1 string, id2 int) error {
	return uc.repo.DeleteByIDs(id1, id2)
}

func (uc *sensorUsecase) UpdateSensorData(data *entity.SensorData) error {
	return uc.repo.Update(data)
}
