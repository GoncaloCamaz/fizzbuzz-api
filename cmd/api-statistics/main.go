/*
Package main implements a command-line tool for collecting and displaying API usage statistics.
*/
package main

import (
	"fizzbuzz-api/internal/api-statistics/app"
	"fizzbuzz-api/pkg/database"
	"os"
)

func main() {
	httpServer := os.Getenv("HTTP_SERVER")
	rpcServer := os.Getenv("GRPC_SERVER")
	serviceName := os.Getenv("SERVICE_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if httpServer == "" {
		httpServer = ":8082"
	}

	if rpcServer == "" {
		rpcServer = ":50051"
	}

	dbConfig := database.NewDBConfig(dbHost, dbPort, dbUser, dbPassword, dbName)
	serviceConfiguration := app.NewStatisticsServiceConfiguration(rpcServer, httpServer, serviceName, dbConfig)

	statisticsService := app.NewStatisticsService(serviceConfiguration)

	if err := statisticsService.StartStatisticsService(); err != nil {
		panic(err)
	}
}
