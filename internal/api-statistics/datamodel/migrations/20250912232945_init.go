package migrations

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type statistic struct {
	bun.BaseModel `bun:"table:statistics,alias:statistics"`
	
	Id                   string    `bun:",pk,type:varchar(64)"`
	RequestKey           string    `bun:",notnull,type:varchar(64)"`
	FirstNumber          int64     `bun:",notnull,default:0"`
	SecondNumber         int64     `bun:",notnull,default:0"`
	FirstReplacementStr  string    `bun:",notnull,type:varchar(64)"`
	SecondReplacementStr string    `bun:",notnull,type:varchar(64)"`
	Limit                int64     `bun:",notnull,default:0"`
	Timestamp            time.Time `bun:",notnull,default:current_timestamp"`
}

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewCreateTable().IfNotExists().Model((*statistic)(nil)).Exec(ctx)

			return err
		})
	}, func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
			_, err := tx.NewDropTable().Model((*statistic)(nil)).Exec(ctx)

			return err
		})
	})
}
