package services

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/domain/entities"
	pb "fizzbuzz-api/internal/api-statistics/handlers/grpc/proto"
)

// StatisticsService provides methods to interact with statistics data
type StatisticsService struct {
	rpcClient pb.StatisticsServiceClient
}

// NewStatisticsService creates a new instance of StatisticsService
func NewStatisticsService(client pb.StatisticsServiceClient) *StatisticsService {
	return &StatisticsService{
		rpcClient: client,
	}
}

// CreateStatisticsRecord creates a new statistics record
func (s *StatisticsService) CreateStatisticsRecord(ctx context.Context, data entities.Statistics) (string, error) {
	protoRequest := &pb.StatisticRequest{
		Limit:           data.Limit,
		Multiplier1:     data.MultipleOne,
		Multiplier2:     data.MultipleTwo,
		ReplacementStr1: data.ReplacementStr1,
		ReplacementStr2: data.ReplacementStr2,
	}

	response, err := s.rpcClient.CreateStatistic(ctx, protoRequest)
	if err != nil {
		return "", err
	}

	requestID := response.RequestId

	return requestID, nil
}
