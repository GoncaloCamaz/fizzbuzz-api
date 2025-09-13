/*
Package http implements the HTTP handler for the Statistics API.
*/
package http

import (
	"fizzbuzz-api/internal/api-statistics/app"
	"fizzbuzz-api/internal/api-statistics/dto/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Service represents an interface for starting a service
type Service interface {
	StartService()
}

// StatisticsHTTPHandler represents a Statistics HTTP handler
type StatisticsHTTPHandler struct {
	echo *echo.Echo
	conf *app.StatisticsServiceConfiguration
	svc  *app.StatisticsService
}

// NewStatisticsHTTPHandler creates a new StatisticsHTTPHandler
func NewStatisticsHTTPHandler(conf *app.StatisticsServiceConfiguration, svc *app.StatisticsService) *StatisticsHTTPHandler {
	return &StatisticsHTTPHandler{
		conf: conf,
		svc:  svc,
	}
}

// StartService starts the Statistics HTTP service
func (h *StatisticsHTTPHandler) StartService() {
	h.echo = echo.New()

	// add http routes
	h.addRoutes()

	// start the server
	if err := h.echo.Start(h.conf.ServiceHTTPPath); err != nil {
		panic(err)
	}
}

func (h *StatisticsHTTPHandler) addRoutes() {
	h.echo.GET("/statistics/get", h.handleGetStatistics)
}

func (h *StatisticsHTTPHandler) handleGetStatistics(c echo.Context) error {
	mostFrequentRequest, count, err := h.svc.Queries.HandleGet(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if mostFrequentRequest == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No statistics found"})
	}

	return c.JSON(http.StatusOK, responses.SerializeMostFrequentRequestResponse(mostFrequentRequest, count))
}
