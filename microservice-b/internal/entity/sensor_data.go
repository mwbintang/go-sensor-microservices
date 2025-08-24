package entity

import "gorm.io/gorm"

type SensorData struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Value      float64
	SensorType string
	ID1        string
	ID2        int
	Timestamp  int64
}
