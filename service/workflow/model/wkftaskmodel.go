package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfTaskModel = (*customWkfTaskModel)(nil)

type (
	// WkfTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfTaskModel.
	WkfTaskModel interface {
		wkfTaskModel
		FindAll(context.Context, int64, int64) ([]*WkfTask, error)
		Count(context.Context) (int64, error)
	}

	customWkfTaskModel struct {
		*defaultWkfTaskModel
	}
)

// NewWkfTaskModel returns a model for the database table.
func NewWkfTaskModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfTaskModel {
	return &customWkfTaskModel{
		defaultWkfTaskModel: newWkfTaskModel(conn, c, opts...),
	}
}

func (m *customWkfTaskModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*WkfTask, error) {
	var resp []*WkfTask
	query := fmt.Sprintf("select * from %s limit ?,?", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, (Current-1)*PageSize, PageSize)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customWkfTaskModel) Count(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query)
	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}
