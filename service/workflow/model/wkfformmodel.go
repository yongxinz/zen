package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfFormModel = (*customWkfFormModel)(nil)

type (
	// WkfFormModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfFormModel.
	WkfFormModel interface {
		wkfFormModel
		FindByTicket(context.Context, int64) ([]*WkfForm, error)
	}

	customWkfFormModel struct {
		*defaultWkfFormModel
	}
)

// NewWkfFormModel returns a model for the database table.
func NewWkfFormModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfFormModel {
	return &customWkfFormModel{
		defaultWkfFormModel: newWkfFormModel(conn, c, opts...),
	}
}

func (m *customWkfFormModel) FindByTicket(ctx context.Context, ticket int64) ([]*WkfForm, error) {
	var resp []*WkfForm
	query := fmt.Sprintf("select * from %s where ticket = ? order by id desc", m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, ticket)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
