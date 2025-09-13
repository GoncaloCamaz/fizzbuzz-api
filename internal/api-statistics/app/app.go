/*
Package app holds utilities to start and manage the application
*/
package app

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel/migrations"
	"fizzbuzz-api/internal/api-statistics/queries"
	"fizzbuzz-api/pkg/database"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

// StatisticsService represents the statistics service
type StatisticsService struct {
	configuration *StatisticsServiceConfiguration
	Queries       *queries.Queries
}

// NewStatisticsService returns a new StatisticsService
func NewStatisticsService(conf *StatisticsServiceConfiguration, db *bun.DB) *StatisticsService {
	return &StatisticsService{
		configuration: conf,
		Queries:       queries.NewQueries(queries.NewStatisticQueriesHandler(db)),
	}
}

func (s *StatisticsService) SetupService() error {
	db := database.NewDB(s.configuration.DB)

	migrator := migrate.NewMigrator(db, migrations.Migrations, migrate.WithMarkAppliedOnSuccess(true))
	if err := migrator.Init(context.Background()); err != nil {
		return err
	}
	if _, err := migrator.Migrate(context.Background()); err != nil {
		return err
	}

	return nil
}
