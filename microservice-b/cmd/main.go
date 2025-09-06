package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/entity"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/repository"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/transport"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/usecase"
	"github.com/mwbintang/go-sensor-microservices/proto"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 1. Initialize Database Connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "Test123!"),
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "sensor_db"),
	)

	log.Printf("Connecting to database: %s@%s:%s", getEnv("DB_USER", ""), getEnv("DB_HOST", ""), getEnv("DB_PORT", ""))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("✅ Successfully connected to database")

	// 2. Auto Migrate the Schema
	err = db.AutoMigrate(&entity.SensorData{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	log.Println("✅ Database schema migrated successfully")

	// 3. Initialize Repository
	sensorRepo := repository.NewMySQLRepository(db)
	log.Println("✅ Repository initialized")

	// 4. Initialize Usecase
	sensorUsecase := usecase.NewSensorUsecase(sensorRepo)
	log.Println("✅ Usecase initialized")

	// 5. Initialize gRPC Server
	grpcServer := grpc.NewServer()
	sensorService := transport.NewSensorServiceServer(sensorUsecase)
	proto.RegisterSensorServiceServer(grpcServer, sensorService)
	log.Println("✅ gRPC server initialized")

	// 6. Start gRPC Server
	grpcLis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on gRPC port: %v", err)
	}

	go func() {
		log.Printf("🚀 gRPC server listening on :50051")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	// Start HTTP server
	httpServer := transport.NewHTTPServer(getEnv("HTTP_PORT", "8082"), sensorUsecase)
	go func() {
		log.Printf("🚀 HTTP server listening on :%s")
		if err := httpServer.Start(); err != nil {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	// 7. Wait for shutdown signal
	waitForShutdown(grpcServer)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func waitForShutdown(grpcServer *grpc.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Press Ctrl+C to shutdown servers...")
	<-sigChan

	log.Println("🛑 Shutting down servers...")
	grpcServer.GracefulStop()
	log.Println("✅ Servers stopped gracefully")
}
