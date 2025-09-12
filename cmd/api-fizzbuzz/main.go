/*
Package main implements the entry point for the FizzBuzz API service.
*/
package main

import (
	"fizzbuzz-api/internal/api-fizzbuzz/app"
	"os"
)

func main() {
	httpServer := os.Getenv("HTTP_SERVER")
	rpcServer := os.Getenv("STATISTICS_API_GRPC")
	serviceName := os.Getenv("SERVICE_NAME")

	if httpServer == "" {
		httpServer = ":8081"
	}

	if rpcServer == "" {
		rpcServer = "api-statistics:50051"
	}

	serviceConfiguration := app.NewFizzBuzzServiceConfiguration(rpcServer, httpServer, serviceName)
	fizzBuzzService := app.NewFizzBuzzService(serviceConfiguration)

	if err := fizzBuzzService.StartFizzBuzzService(); err != nil {
		panic(err)
	}
}
