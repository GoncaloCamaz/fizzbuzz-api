package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Any("/fizzbuzz/*", func(c echo.Context) error {
		target := "http://api-fizzbuzz:8081" + c.Request().URL.Path
		return proxyRequest(c, target)
	})

	e.Any("/statistics/*", func(c echo.Context) error {
		target := "http://api-statistics:8082" + c.Request().URL.Path
		return proxyRequest(c, target)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

// proxyRequest forwards the request to the correct service
func proxyRequest(c echo.Context, target string) error {
	req, err := http.NewRequest(c.Request().Method, target, c.Request().Body)
	if err != nil {
		return err
	}

	for k, v := range c.Request().Header {
		req.Header.Set(k, v[0])
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return c.Stream(resp.StatusCode, resp.Header.Get("Content-Type"), resp.Body)
}
