/*
Package app holds utilities to start and manage the application
*/
package app

import (
	"fizzbuzz-api/internal/api-statistics/handlers/grpc"
	"fizzbuzz-api/internal/api-statistics/handlers/http"
	"fmt"
)

// StatisticsService represents the statistics service
type StatisticsService struct {
	configuration *StatisticsServiceConfiguration
}

// NewStatisticsService returns a new StatisticsService
func NewStatisticsService(conf *StatisticsServiceConfiguration) *StatisticsService {
	return &StatisticsService{
		configuration: conf,
	}
}

func (s *StatisticsService) StartStatisticsService() error {
	fmt.Println("Starting statistics service", s.configuration)

	httpHandler := http.NewStatisticsHTTPHandler(s.configuration.ServiceHTTPPath)
	httpHandler.StartService()

	rpcHandler := grpc.NewStatisticsGRPCHandler(s.configuration.ServiceRPCPath)
	rpcHandler.StartGRPCService()

	return nil
}
