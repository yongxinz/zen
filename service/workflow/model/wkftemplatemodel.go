package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfTemplateModel = (*customWkfTemplateModel)(nil)

type (
	// WkfTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfTemplateModel.
	WkfTemplateModel interface {
		wkfTemplateModel
		FindAll(context.Context, int64, int64) ([]*WkfTemplate, error)
		Count(context.Context) (int64, error)
	}

	customWkfTemplateModel struct {
		*defaultWkfTemplateModel
	}
)

// NewWkfTemplateModel returns a model for the database table.
func NewWkfTemplateModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfTemplateModel {
	return &customWkfTemplateModel{
		defaultWkfTemplateModel: newWkfTemplateModel(conn, c, opts...),
	}
}

func (m *customWkfTemplateModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*WkfTemplate, error) {
	var resp []*WkfTemplate
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

func (m *customWkfTemplateModel) Count(ctx context.Context) (int64, error) {
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
