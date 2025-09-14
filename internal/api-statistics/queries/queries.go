/*
Package queries provides the query implementations for the Statistics service.
*/
package queries

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel"
)

// Queries represents the queries for the Statistics service
type Queries struct {
	StatisticQueries
}

// NewQueries creates a new Queries instance
func NewQueries(sq StatisticQueries) *Queries {
	return &Queries{
		StatisticQueries: sq,
	}
}

// StatisticQueries defines the interface for statistic queries
type StatisticQueries interface {
	HandleGet(ctx context.Context) (*datamodel.Statistic, int, error)
}
