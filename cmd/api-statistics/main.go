/*
Package main implements a command-line tool for collecting and displaying API usage statistics.
*/
package main

import (
	"fizzbuzz-api/internal/api-statistics/app"
	"os"
)

func main() {
	httpServer := os.Getenv("HTTP_SERVER")
	rpcServer := os.Getenv("GRPC_SERVER")
	serviceName := os.Getenv("SERVICE_NAME")

	if httpServer == "" {
		httpServer = ":8082"
	}

	if rpcServer == "" {
		rpcServer = ":50051"
	}

	serviceConfiguration := app.NewStatisticsServiceConfiguration(rpcServer, httpServer, serviceName)
	statisticsService := app.NewStatisticsService(serviceConfiguration)

	if err := statisticsService.StartStatisticsService(); err != nil {
		panic(err)
	}
}
