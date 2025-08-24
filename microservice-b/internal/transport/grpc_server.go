package transport

import (
	"context"
	"log"

	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/entity"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/usecase"
	"github.com/mwbintang/go-sensor-microservices/proto"
)

type SensorServiceServer struct {
	proto.UnimplementedSensorServiceServer
	usecase usecase.SensorUsecase
}

func NewSensorServiceServer(usecase usecase.SensorUsecase) *SensorServiceServer {
	return &SensorServiceServer{usecase: usecase}
}

func (s *SensorServiceServer) SendSensorData(ctx context.Context, req *proto.SensorData) (*proto.SensorResponse, error) {
	log.Printf("Received sensor data: %v", req)

	// Convert proto SensorData to entity SensorData
	data := &entity.SensorData{
		ID:         req.Id,
		Value:      req.Value,
		SensorType: req.SensorType,
		ID1:        req.Id1,
		ID2:        int(req.Id2),
		Timestamp:  req.Timestamp,
	}

	err := s.usecase.CreateSensorData(data)
	if err != nil {
		return &proto.SensorResponse{Success: false, Message: err.Error()}, err
	}

	return &proto.SensorResponse{Success: true, Message: "Data stored successfully"}, nil
}
