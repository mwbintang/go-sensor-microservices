package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/usecase"
)

type SensorHandler struct {
	usecase usecase.SensorUsecase
}

func NewSensorHandler(usecase usecase.SensorUsecase) *SensorHandler {
	return &SensorHandler{usecase: usecase}
}

// Get sensor data by IDs
func (h *SensorHandler) GetDataByIDs(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2Str := c.QueryParam("id2")

	id2, err := strconv.Atoi(id2Str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID2"})
	}

	data, err := h.usecase.FindByIDs(id1, id2)
	fmt.Println(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Endpoint working",
		"id1":     id1,
		"id2":     id2Str,
	})
}

// Get sensor data by time range
func (h *SensorHandler) GetDataByTimeRange(c echo.Context) error {
	start := c.QueryParam("start")
	end := c.QueryParam("end")

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Time range endpoint",
		"start":   start,
		"end":     end,
	})
}

// Health check endpoint
func (h *SensorHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
