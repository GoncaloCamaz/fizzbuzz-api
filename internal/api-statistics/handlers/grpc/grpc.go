/*
Package grpc provides gRPC handlers for the API statistics service.
*/
package grpc

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/app"
	"fizzbuzz-api/internal/api-statistics/domain/entities"
	"net"

	pb "fizzbuzz-api/internal/api-statistics/handlers/grpc/proto"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

// Server implements the gRPC server for the statistics service.
type Server struct {
	pb.StatisticsServiceServer
	svc *app.StatisticsService
}

// ServiceGRPC defines the interface for starting the gRPC service.
type ServiceGRPC interface {
	StartGRPCService()
}

// StatisticsGRPCHandler represents a gRPC handler for the statistics service.
type StatisticsGRPCHandler struct {
	svc  *app.StatisticsService
	conf *app.StatisticsServiceConfiguration
}

// NewStatisticsGRPCHandler creates a new instance of StatisticsGRPCHandler.
func NewStatisticsGRPCHandler(conf *app.StatisticsServiceConfiguration, svc *app.StatisticsService) *StatisticsGRPCHandler {
	return &StatisticsGRPCHandler{
		conf: conf,
		svc:  svc,
	}
}

// StartGRPCService starts the gRPC service.
func (h *StatisticsGRPCHandler) StartGRPCService() {
	lis, err := net.Listen("tcp", h.conf.ServiceRPCPath)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterStatisticsServiceServer(s, &Server{
		svc: h.svc,
	})

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *Server) CreateStatistic(ctx context.Context, req *pb.StatisticRequest) (*pb.StatisticResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	statisticEntity := entities.Statistics{
		ID:              id.String(),
		MultipleOne:     req.Multiplier1,
		MultipleTwo:     req.Multiplier2,
		ReplacementStr1: req.ReplacementStr1,
		ReplacementStr2: req.ReplacementStr2,
		Limit:           req.Limit,
	}

	err = s.svc.Service.CreateStatistics(ctx, &statisticEntity)
	if err != nil {
		return nil, err
	}

	return &pb.StatisticResponse{
		RequestId: id.String(),
	}, nil
}
