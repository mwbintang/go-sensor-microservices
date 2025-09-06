package repository

import "github.com/mwbintang/go-sensor-microservices/microservice-b/internal/entity"

type SensorRepository interface {
	Create(data *entity.SensorData) error
	FindByIDs(id1 string, id2 int) (*entity.SensorData, error)
	FindByTimeRange(start, end int64) ([]entity.SensorData, error)
	DeleteByIDs(id1 string, id2 int) error
	Update(data *entity.SensorData) error
	FindPaginated(page, pageSize int) ([]entity.SensorData, int64, error)
}
