/*
Package datamodel defines the data structures used for API statistics.
*/
package datamodel

import (
	"time"

	"github.com/uptrace/bun"
)

type Statistic struct {
	bun.BaseModel `bun:"table:statistics,alias:statistics"`

	Id              string    `bun:",pk,type:varchar(64)"`
	RequestKey      string    `bun:",notnull,type:varchar(64)"`
	MultipleOne     int64     `bun:",notnull,default:0"`
	MultipleTwo     int64     `bun:",notnull,default:0"`
	ReplacementStr1 string    `bun:",notnull,type:varchar(64)"`
	ReplacementStr2 string    `bun:",notnull,type:varchar(64)"`
	Limit           int64     `bun:",notnull,default:0"`
	Timestamp       time.Time `bun:",notnull,default:current_timestamp"`
}
