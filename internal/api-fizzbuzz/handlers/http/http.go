/*
Package http implements the HTTP handler for the FizzBuzz API.
*/
package http

import (
	"context"
	"fizzbuzz-api/internal/api-fizzbuzz/app"
	"fizzbuzz-api/internal/api-fizzbuzz/dto/requests"
	"fizzbuzz-api/internal/api-fizzbuzz/dto/responses"
	"fizzbuzz-api/internal/api-statistics/domain/entities"
	"fizzbuzz-api/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Service represents an interface for starting a service
type Service interface {
	StartService()
}

// FizzBuzzHTTPHandler represents a FizzBuzz HTTP handler
type FizzBuzzHTTPHandler struct {
	echo *echo.Echo
	conf *app.FizzBuzzServiceConfiguration
	svc  *app.FizzBuzzService
}

// NewFizzBuzzHTTPHandler creates a new FizzBuzzHTTPHandler
func NewFizzBuzzHTTPHandler(conf *app.FizzBuzzServiceConfiguration, svc *app.FizzBuzzService) *FizzBuzzHTTPHandler {
	return &FizzBuzzHTTPHandler{
		conf: conf,
		svc:  svc,
	}
}

// StartService starts the FizzBuzz HTTP service
func (h *FizzBuzzHTTPHandler) StartService() {
	h.echo = echo.New()

	// add http routes
	h.addRoutes()

	// start the server
	if err := h.echo.Start(h.conf.ServiceHTTPPath); err != nil {
		panic(err)
	}
}

func (h *FizzBuzzHTTPHandler) addRoutes() {
	h.echo.POST("/fizzbuzz/get", h.handleFizzBuzz)
}

func (h *FizzBuzzHTTPHandler) handleFizzBuzz(c echo.Context) error {
	fizzBuzzDTO := new(requests.FizzBuzzRequest)
	if err := c.Bind(fizzBuzzDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := fizzBuzzDTO.ValidateFizzBuzzParams(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fizzBuzzResult := utils.FizzBuzz(fizzBuzzDTO.Limit, fizzBuzzDTO.FirstNumber, fizzBuzzDTO.SecondNumber,
		fizzBuzzDTO.FirstReplacementStr, fizzBuzzDTO.SecondReplacementStr)

	go func() {
		_, err := h.svc.Service.CreateStatisticsRecord(context.WithoutCancel(c.Request().Context()), entities.Statistics{
			FirstNumber:          fizzBuzzDTO.FirstNumber,
			SecondNumber:         fizzBuzzDTO.SecondNumber,
			FirstReplacementStr:  fizzBuzzDTO.FirstReplacementStr,
			SecondReplacementStr: fizzBuzzDTO.SecondReplacementStr,
			Limit:                fizzBuzzDTO.Limit,
		})
		if err != nil {
			log.Err(err).Msg("Failed to create statistics record")
		}
	}()

	return c.JSON(http.StatusOK, responses.SerializeFizzBuzzResponse(fizzBuzzResult))
}
