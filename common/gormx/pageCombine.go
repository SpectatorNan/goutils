package gormx

import (
	"context"
	"github.com/SpectatorNan/gorm-zero/gormc"
	"gorm.io/gorm"
	"goutils/common/pagex"
)

var tableSortDesc = "descend"
var tableSortAsc = "ascend"

func SetTableSortAsc(key string) {
	tableSortAsc = key
}
func SetTableSortDesc(key string) {
	tableSortDesc = key
}

func FindPageList[T any](ctx context.Context, cc gormc.CachedConn, page *pagex.ListReq, orderKey, sort string, orderKeys map[string]string, fn func(conn *gorm.DB) *gorm.DB) ([]T, int64, error) {
	var res []T
	var count int64
	err := cc.ExecNoCacheCtx(ctx, func(conn *gorm.DB) error {
		return fn(conn).Count(&count).Error
	})
	if err != nil {
		return nil, 0, err
	}
	err = cc.QueryNoCacheCtx(ctx, &res, func(conn *gorm.DB, v interface{}) error {
		db := fn(conn).Scopes(Paginate(page))
		if orderStr, ok := orderKeys[orderKey]; ok {
			if sort == tableSortDesc {
				db = db.Order(orderStr + " desc")
			} else {
				db = db.Order(orderStr + " asc")
			}
		}
		return db.Find(v).Error
	})
	if err != nil {
		return nil, 0, err
	}
	return res, count, nil
}
