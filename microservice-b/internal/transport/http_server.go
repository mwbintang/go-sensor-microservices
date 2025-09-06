package transport

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/handler"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/usecase"
)

type HTTPServer struct {
	echo *echo.Echo
	port string
}

func NewHTTPServer(port string, sensorUsecase usecase.SensorUsecase) *HTTPServer {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Initialize handler
	sensorHandler := handler.NewSensorHandler(sensorUsecase)

	// Routes
	e.GET("/health", sensorHandler.HealthCheck)
	e.GET("/data", sensorHandler.GetDataPaginated)
	e.DELETE("/data", sensorHandler.DeleteData)
	e.PUT("/data", sensorHandler.EditData)

	return &HTTPServer{
		echo: e,
		port: port,
	}
}

func (s *HTTPServer) Start() error {
	return s.echo.Start(":" + s.port)
}

func (s *HTTPServer) Shutdown() error {
	return s.echo.Close()
}
