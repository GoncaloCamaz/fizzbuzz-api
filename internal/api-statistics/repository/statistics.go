package repository

import (
	"context"
	"fizzbuzz-api/internal/api-statistics/datamodel"

	"github.com/uptrace/bun"
)

// CreateStatistics creates a new statistics record
func (p *PGRepo) CreateStatistics(ctx context.Context, db bun.IDB, st *datamodel.Statistic) error {
	_, err := db.NewInsert().Model(st).Exec(ctx)
	return err
}

// GetMostPerformedRequestKey retrieves the most performed request key and its count
func (p *PGRepo) GetMostPerformedRequestKey(ctx context.Context, db bun.IDB) (string, int) {
	var result struct {
		RequestKey string `bun:"request_key"`
		Count      int    `bun:"count"`
	}

	err := db.NewSelect().
		Model((*datamodel.Statistic)(nil)).
		Column("request_key").
		ColumnExpr("COUNT(*) AS count").
		Group("request_key").
		Order("count DESC").
		Limit(1).
		Scan(ctx, &result)
	if err != nil {
		return "", 0
	}

	return result.RequestKey, result.Count
}

// GetStatisticByRequestKey retrieves a statistics record by request key
func (p *PGRepo) GetStatisticByRequestKey(ctx context.Context, db bun.IDB, requestKey string) (*datamodel.Statistic, error) {
	var stat datamodel.Statistic
	err := db.NewSelect().
		Model(&stat).
		Where("request_key = ?", requestKey).
		Order("created_at DESC").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return &stat, nil
}
