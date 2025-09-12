/*
Package http implements the HTTP handler for the FizzBuzz API.
*/
package http

import (
	"fizzbuzz-api/internal/api-fizzbuzz/dto/requests"
	"fizzbuzz-api/internal/api-fizzbuzz/dto/responses"
	"fizzbuzz-api/pkg/utils"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Service represents an interface for starting a service
type Service interface {
	StartService()
}

// FizzBuzzHTTPHandler represents a FizzBuzz HTTP handler
type FizzBuzzHTTPHandler struct {
	echo *echo.Echo
	port string
}

// NewFizzBuzzHTTPHandler creates a new FizzBuzzHTTPHandler
func NewFizzBuzzHTTPHandler(port string) *FizzBuzzHTTPHandler {
	return &FizzBuzzHTTPHandler{
		port: port,
	}
}

// StartService starts the FizzBuzz HTTP service
func (h *FizzBuzzHTTPHandler) StartService() {
	fmt.Println(h.port)
	h.echo = echo.New()

	// add http routes
	h.addRoutes()

	// start the server
	if err := h.echo.Start(h.port); err != nil {
		panic(err)
	}
}

func (h *FizzBuzzHTTPHandler) addRoutes() {
	h.echo.POST("/fizzbuzz", h.handleFizzBuzz)
}

func (h *FizzBuzzHTTPHandler) handleFizzBuzz(c echo.Context) error {
	fizzBuzzDTO := new(requests.FizzBuzzRequest)
	if err := c.Bind(fizzBuzzDTO); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := fizzBuzzDTO.ValidateFizzBuzzParams(); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	fizzBuzzResult := utils.FizzBuzz(fizzBuzzDTO.Limit, fizzBuzzDTO.Multiple1, fizzBuzzDTO.Multiple2,
		fizzBuzzDTO.ReplacementString1, fizzBuzzDTO.ReplacementString2)

	return c.JSON(http.StatusOK, responses.SerializeFizzBuzzResponse(fizzBuzzResult))
}
