package queries

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel"
	"fizzbuzz-api/internal/api-statistics/repository"
	"fmt"

	"github.com/uptrace/bun"
)

type StatisticQueriesHandler struct {
	repo *repository.PGRepo
	db   bun.IDB
}

func NewStatisticQueriesHandler(db bun.IDB) *StatisticQueriesHandler {
	return &StatisticQueriesHandler{
		repo: repository.NewPGRepo(),
		db:   db,
	}
}

func (s StatisticQueriesHandler) HandleGet(ctx context.Context) (*datamodel.Statistic, int, error) {
	mostPerformedRequestKey, count := s.repo.GetMostPerformedRequestKey(ctx, s.db)
	if mostPerformedRequestKey == "" {
		return &datamodel.Statistic{}, 0, fmt.Errorf("no most performed request key")
	}

	st, err := s.repo.GetStatisticByRequestKey(ctx, s.db, mostPerformedRequestKey)
	if err != nil {
		return nil, 0, err
	}

	if st == nil {
		return &datamodel.Statistic{}, 0, fmt.Errorf("no statistics found for the most performed request key")
	}

	return st, count, nil
}
