package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/entity"
	"github.com/mwbintang/go-sensor-microservices/microservice-b/internal/usecase"
)

type SensorHandler struct {
	usecase usecase.SensorUsecase
}

func NewSensorHandler(usecase usecase.SensorUsecase) *SensorHandler {
	return &SensorHandler{usecase: usecase}
}

// GET: retrieve data by IDs, by time range, or both (with optional pagination)
func (h *SensorHandler) GetDataPaginated(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	if pageSize < 1 {
		pageSize = 10
	}

	data, total, err := h.usecase.FindPaginated(page, pageSize)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":       data,
		"page":       page,
		"page_size":  pageSize,
		"total":      total,
		"total_page": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// DELETE: delete data by IDs
func (h *SensorHandler) DeleteData(c echo.Context) error {
	id1 := c.QueryParam("id1")
	id2Str := c.QueryParam("id2")

	if id1 == "" || id2Str == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "id1 and id2 are required"})
	}

	id2, err := strconv.Atoi(id2Str)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid id2"})
	}

	if err := h.usecase.DeleteByIDs(id1, id2); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Data deleted successfully"})
}

// PUT: update data by IDs
func (h *SensorHandler) EditData(c echo.Context) error {
	var req entity.SensorData
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if err := h.usecase.UpdateSensorData(&req); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Data updated successfully"})
}

func (h *SensorHandler) HealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
}
