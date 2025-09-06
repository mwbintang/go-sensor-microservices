package repository

import (
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/entity"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) *MySQLRepository {
	return &MySQLRepository{db: db}
}

func (r *MySQLRepository) Create(sensorData *entity.SensorData) error {
	return r.db.Create(sensorData).Error
}

func (r *MySQLRepository) FindByIDs(id1 string, id2 int) (*entity.SensorData, error) {
	var data entity.SensorData
	if err := r.db.Where("id1 = ? AND id2 = ?", id1, id2).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *MySQLRepository) FindByTimeRange(start, end int64) ([]entity.SensorData, error) {
	var data []entity.SensorData
	if err := r.db.Where("timestamp BETWEEN ? AND ?", start, end).Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (r *MySQLRepository) DeleteByIDs(id1 string, id2 int) error {
	return r.db.Where("id1 = ? AND id2 = ?", id1, id2).Delete(&entity.SensorData{}).Error
}

func (r *MySQLRepository) Update(sensorData *entity.SensorData) error {
	return r.db.Save(sensorData).Error
}

func (r *MySQLRepository) FindPaginated(page, pageSize int) ([]entity.SensorData, int64, error) {
	var data []entity.SensorData
	var total int64

	r.db.Model(&entity.SensorData{}).Count(&total) // get total rows

	offset := (page - 1) * pageSize
	result := r.db.Limit(pageSize).Offset(offset).Find(&data)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return data, total, nil
}
