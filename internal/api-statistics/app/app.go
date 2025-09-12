/*
Package app holds utilities to start and manage the application
*/
package app

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel/migrations"
	"fizzbuzz-api/internal/api-statistics/handlers/grpc"
	"fizzbuzz-api/internal/api-statistics/handlers/http"
	"fizzbuzz-api/pkg/database"
	"fmt"

	"github.com/uptrace/bun/migrate"
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

	err := SetupService(s.configuration)
	if err != nil {
		return err
	}

	httpHandler := http.NewStatisticsHTTPHandler(s.configuration.ServiceHTTPPath)
	httpHandler.StartService()

	rpcHandler := grpc.NewStatisticsGRPCHandler(s.configuration.ServiceRPCPath)
	rpcHandler.StartGRPCService()

	return nil
}

func SetupService(conf *StatisticsServiceConfiguration) error {
	db := database.NewDB(conf.DB)

	migrator := migrate.NewMigrator(db, migrations.Migrations, migrate.WithMarkAppliedOnSuccess(true))
	if err := migrator.Init(context.Background()); err != nil {
		return err
	}
	if _, err := migrator.Migrate(context.Background()); err != nil {
		return err
	}

	return nil
}
