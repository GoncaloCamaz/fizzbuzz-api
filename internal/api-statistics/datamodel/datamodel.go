/*
Package datamodel defines the data structures used for API statistics.
*/
package datamodel

import (
	"time"

	"github.com/uptrace/bun"
)

// Statistic represents the statistics data model for database storage
type Statistic struct {
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
