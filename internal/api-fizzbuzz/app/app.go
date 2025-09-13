/*
Package app holds utilities to start and manage the application
*/
package app

import (
	"fizzbuzz-api/internal/api-fizzbuzz/domain/services"
	pb "fizzbuzz-api/internal/api-statistics/handlers/grpc/proto"
)

// FizzBuzzService represents the fizzbuzz service
type FizzBuzzService struct {
	configuration *FizzBuzzServiceConfiguration
	Service       *services.StatisticsService
}

// NewFizzBuzzService returns a new FizzBuzzService
func NewFizzBuzzService(conf *FizzBuzzServiceConfiguration, client pb.StatisticsServiceClient) *FizzBuzzService {
	return &FizzBuzzService{
		configuration: conf,
		Service:       services.NewStatisticsService(client),
	}
}
