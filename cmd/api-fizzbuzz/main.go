/*
Package main implements the entry point for the FizzBuzz API service.
*/
package main

import (
	"fizzbuzz-api/internal/api-fizzbuzz/app"
	"fizzbuzz-api/internal/api-fizzbuzz/handlers/http"
	pb "fizzbuzz-api/internal/api-statistics/handlers/grpc/proto"
	"fmt"
	"os"

	"google.golang.org/grpc"
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
	conn, err := grpc.NewClient(rpcServer, grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
	}

	rpcClient := pb.NewStatisticsServiceClient(conn)
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(conn)

	statisticsService := app.NewFizzBuzzService(serviceConfiguration, rpcClient)

	httpHandler := http.NewFizzBuzzHTTPHandler(serviceConfiguration, statisticsService)
	httpHandler.StartService()
}
