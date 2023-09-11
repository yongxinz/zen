package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfCirculationModel = (*customWkfCirculationModel)(nil)

type (
	// WkfCirculationModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfCirculationModel.
	WkfCirculationModel interface {
		wkfCirculationModel
		FindByTicket(context.Context, int64) ([]*WkfCirculation, error)
	}

	customWkfCirculationModel struct {
		*defaultWkfCirculationModel
	}
)

// NewWkfCirculationModel returns a model for the database table.
func NewWkfCirculationModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfCirculationModel {
	return &customWkfCirculationModel{
		defaultWkfCirculationModel: newWkfCirculationModel(conn, c, opts...),
	}
}

func (m *customWkfCirculationModel) FindByTicket(ctx context.Context, ticket int64) ([]*WkfCirculation, error) {
	var resp []*WkfCirculation
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
