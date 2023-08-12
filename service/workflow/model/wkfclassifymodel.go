package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfClassifyModel = (*customWkfClassifyModel)(nil)

type (
	// WkfClassifyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfClassifyModel.
	WkfClassifyModel interface {
		wkfClassifyModel
		FindAll(context.Context, int64, int64) ([]*WkfClassify, error)
		Count(context.Context) (int64, error)
	}

	customWkfClassifyModel struct {
		*defaultWkfClassifyModel
	}
)

// NewWkfClassifyModel returns a model for the database table.
func NewWkfClassifyModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfClassifyModel {
	return &customWkfClassifyModel{
		defaultWkfClassifyModel: newWkfClassifyModel(conn, c, opts...),
	}
}

func (m *customWkfClassifyModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*WkfClassify, error) {
	var resp []*WkfClassify
	query := "select wkf_classify.* from wkf_classify limit ?,?"
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

func (m *customWkfClassifyModel) Count(ctx context.Context) (int64, error) {
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
