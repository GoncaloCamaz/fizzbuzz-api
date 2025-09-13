package queries

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel"
)

type Queries struct {
	StatisticQueries
}

func NewQueries(sq StatisticQueries) *Queries {
	return &Queries{
		StatisticQueries: sq,
	}
}

type StatisticQueries interface {
	HandleGet(ctx context.Context) (*datamodel.Statistic, int, error)
}
