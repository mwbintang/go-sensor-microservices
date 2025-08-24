package main

import (
	"log"
	"os"
	"strconv"

	"github.com/mwbintang/go-sensor-microservices/microservice-a/internal/generator"
	"github.com/mwbintang/go-sensor-microservices/microservice-a/internal/server"
)

func main() {
	// Get configuration from environment variables
	microserviceBURL := getEnv("MICROSERVICE_B_URL", "localhost:50051")
	sensorType := getEnv("SENSOR_TYPE", "temperature") // Fixed sensor type per instance
	initialIntervalMs := getEnvAsInt("INITIAL_INTERVAL_MS", 1000)
	httpPort := getEnv("HTTP_PORT", "8081")

	// Create sensor generator
	generator := generator.NewSensorGenerator(microserviceBURL, sensorType, initialIntervalMs)

	// Start generating data
	generator.Start()
	defer generator.Stop()

	// Create and start HTTP server
	httpServer := server.NewServer(httpPort, generator)

	log.Printf("Microservice A (%s) starting on port %s", sensorType, httpPort)
	log.Printf("Sending data to Microservice B at: %s", microserviceBURL)
	log.Printf("Initial interval: %dms", initialIntervalMs)

	if err := httpServer.Start(); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
