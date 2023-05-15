package gormx

import (
	"fmt"
	"gorm.io/gorm"
	"goutils/common/pagex"
	"time"
)

func CreateTimeRange(startTime, endTime *time.Time) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if startTime != nil {
			db = db.Where("created_at >= ?", *startTime)
		}
		if endTime != nil {
			db = db.Where("created_at <= ?", *endTime)
		}
		return db
	}
}

func Paginate(page *pagex.ListReq) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == nil {
			return db
		}
		return db.Offset(page.Offset()).Limit(page.Limit())
	}
}

func AvgRaw(column, alias string) string {
	return fmt.Sprintf("IFNULL(avg(%s),0) as %s", column, alias)
}

func CaseWhenNull(column, alias string, defaultValue interface{}) string {
	var fmtStr = "case WHEN avg( %s )  IS NULL THEN %v ELSE avg( %s ) END %s"
	if _, ok := defaultValue.(int); ok {
		fmtStr = "case WHEN avg( %s )  IS NULL THEN %d ELSE avg( %s ) END %s"
	} else if _, ok := defaultValue.(float64); ok {
		fmtStr = "case WHEN avg( %s )  IS NULL THEN %.2f ELSE avg( %s ) END %s"
	}

	return fmt.Sprintf(fmtStr, column, defaultValue, column, alias)
}
