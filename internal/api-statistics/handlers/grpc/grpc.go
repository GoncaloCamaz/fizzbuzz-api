/*
Package grpc provides gRPC handlers for the API statistics service.
*/
package grpc

import (
	"context"
	"net"

	pb "fizzbuzz-api/internal/api-statistics/handlers/grpc/proto"

	"google.golang.org/grpc"
)

// Server implements the gRPC server for the statistics service.
type Server struct {
	pb.StatisticsServiceServer
}

// ServiceGRPC defines the interface for starting the gRPC service.
type ServiceGRPC interface {
	StartGRPCService()
}

// StatisticsGRPCHandler represents a gRPC handler for the statistics service.
type StatisticsGRPCHandler struct {
	Address string
}

// NewStatisticsGRPCHandler creates a new instance of StatisticsGRPCHandler.
func NewStatisticsGRPCHandler(address string) *StatisticsGRPCHandler {
	return &StatisticsGRPCHandler{
		Address: address,
	}
}

// StartGRPCService starts the gRPC service.
func (h *StatisticsGRPCHandler) StartGRPCService() {
	lis, err := net.Listen("tcp", h.Address)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	pb.RegisterStatisticsServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}

func (s *Server) CreateStatistic(ctx context.Context, req *pb.StatisticRequest) (*pb.StatisticResponse, error) {
	return nil, nil
}
