package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfExecutionModel = (*customWkfExecutionModel)(nil)

type (
	// WkfExecutionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfExecutionModel.
	WkfExecutionModel interface {
		wkfExecutionModel
	}

	customWkfExecutionModel struct {
		*defaultWkfExecutionModel
	}
)

// NewWkfExecutionModel returns a model for the database table.
func NewWkfExecutionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfExecutionModel {
	return &customWkfExecutionModel{
		defaultWkfExecutionModel: newWkfExecutionModel(conn, c, opts...),
	}
}
