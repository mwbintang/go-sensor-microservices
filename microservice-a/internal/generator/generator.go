package generator

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/mwbintang/go-sensor-microservices/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type SensorGenerator struct {
	client     proto.SensorServiceClient
	conn       *grpc.ClientConn
	sensorType string
	interval   time.Duration
	ticker     *time.Ticker
	stopChan   chan bool
}

func NewSensorGenerator(microserviceBURL, sensorType string, intervalMs int) *SensorGenerator {
	// Connect to Microservice B
	conn, err := grpc.Dial(microserviceBURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to Microservice B: %v", err)
	}

	client := proto.NewSensorServiceClient(conn)

	return &SensorGenerator{
		client:     client,
		conn:       conn,
		sensorType: sensorType,
		interval:   time.Duration(intervalMs) * time.Millisecond,
		stopChan:   make(chan bool),
	}
}

func (g *SensorGenerator) Start() {
	g.ticker = time.NewTicker(g.interval)
	log.Printf("Starting %s sensor data generation every %v", g.sensorType, g.interval)

	go func() {
		for {
			select {
			case <-g.ticker.C:
				g.generateAndSendData()
			case <-g.stopChan:
				g.ticker.Stop()
				g.conn.Close()
				return
			}
		}
	}()
}

func (g *SensorGenerator) Stop() {
	g.stopChan <- true
	log.Printf("Stopped %s sensor data generation", g.sensorType)
}

func (g *SensorGenerator) SetInterval(intervalMs int) {
	g.interval = time.Duration(intervalMs) * time.Millisecond
	if g.ticker != nil {
		g.ticker.Reset(g.interval)
	}
	log.Printf("Changed %s sensor interval to %v", g.sensorType, g.interval)
}

func (g *SensorGenerator) generateAndSendData() {
	data := &proto.SensorData{
		Id:         generateID(),
		Value:      float64(rand.Float32() * 100),
		SensorType: g.sensorType,
		Id1:        generateRandomID1(),
		Id2:        int32(rand.Intn(1000)),
		Timestamp:  time.Now().Unix(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	response, err := g.client.SendSensorData(ctx, data)
	if err != nil {
		log.Printf("Failed to send data to Microservice B: %v", err)
		return
	}

	if response.Success {
		log.Printf("Sent %s data: ID1=%s, ID2=%d, Value=%.2f",
			g.sensorType, data.Id1, data.Id2, data.Value)
	} else {
		log.Printf("Microservice B error: %s", response.Message)
	}
}

// Helper functions
func generateID() string {
	return fmt.Sprintf("sensor-%d", time.Now().UnixNano())
}

func generateRandomID1() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 3)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Add this method here, inside generator package
func (g *SensorGenerator) GetSensorType() string {
	return g.sensorType
}
