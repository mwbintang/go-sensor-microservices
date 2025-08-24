package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mwbintang/go-sensor-microservices/microservice-a/internal/generator"
)

type Server struct {
	echo      *echo.Echo
	generator *generator.SensorGenerator
	port      string
}

func NewServer(port string, generator *generator.SensorGenerator) *Server {
	e := echo.New()
	s := &Server{
		echo:      e,
		generator: generator,
		port:      port,
	}

	// Routes
	e.POST("/frequency", s.setFrequency)
	e.GET("/health", s.healthCheck)

	return s
}

func (s *Server) Start() error {
	return s.echo.Start(":" + s.port)
}

func (s *Server) setFrequency(c echo.Context) error {
	type FrequencyRequest struct {
		IntervalMs int `json:"interval_ms"`
	}

	var req FrequencyRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if req.IntervalMs < 100 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Interval must be at least 100ms"})
	}

	s.generator.SetInterval(req.IntervalMs)

	return c.JSON(http.StatusOK, map[string]string{
		"message":     "Frequency updated successfully",
		"interval":    strconv.Itoa(req.IntervalMs) + "ms",
		"sensor_type": s.generator.GetSensorType(),
	})
}

func (s *Server) healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"status":      "healthy",
		"service":     "microservice-a",
		"sensor_type": s.generator.GetSensorType(),
	})
}
