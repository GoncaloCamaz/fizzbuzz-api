/*
Package services provides business logic for handling statistics in the FizzBuzz API.
*/
package services

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel"
	"fizzbuzz-api/internal/api-statistics/domain/entities"
	"fizzbuzz-api/internal/api-statistics/repository"
	"fmt"
	"time"

	"github.com/uptrace/bun"
)

// StatisticsService provides methods to handle statistics operations.
type StatisticsService struct {
	db   bun.IDB
	repo *repository.PGRepo
}

// NewStatisticsService creates a new instance of StatisticsService.
func NewStatisticsService(db bun.IDB) *StatisticsService {
	return &StatisticsService{
		db:   db,
		repo: repository.NewPGRepo(),
	}
}

// CreateStatistics creates a new statistics record in the database.
func (s *StatisticsService) CreateStatistics(ctx context.Context, st *entities.Statistics) error {
	if st == nil {
		return nil
	}

	requestKey := fmt.Sprintf("%d-%d-%s-%s-%d", st.MultipleOne, st.MultipleTwo, st.ReplacementStr1, st.ReplacementStr2, st.Limit)

	// map entity to datamodel
	dm := &datamodel.Statistic{
		Id:              st.ID,
		RequestKey:      requestKey,
		MultipleOne:     st.MultipleOne,
		MultipleTwo:     st.MultipleTwo,
		ReplacementStr1: st.ReplacementStr1,
		ReplacementStr2: st.ReplacementStr2,
		Limit:           st.Limit,
		Timestamp:       time.Now(),
	}

	return s.repo.CreateStatistics(ctx, s.db, dm)
}
