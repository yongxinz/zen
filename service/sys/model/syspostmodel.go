package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysPostModel = (*customSysPostModel)(nil)

type (
	// SysPostModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysPostModel.
	SysPostModel interface {
		sysPostModel
		FindAll(context.Context, int64, int64) ([]*SysPost, error)
		Count(context.Context) (int64, error)
		DeleteMulti(context.Context, []int64) error
	}

	customSysPostModel struct {
		*defaultSysPostModel
	}
)

// NewSysPostModel returns a model for the database table.
func NewSysPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysPostModel {
	return &customSysPostModel{
		defaultSysPostModel: newSysPostModel(conn, c, opts...),
	}
}

func (m *customSysPostModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysPost, error) {
	var resp []*SysPost
	query := "select sys_post.* from sys_post limit ?,?"
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

func (m *customSysPostModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSysPostModel) DeleteMulti(ctx context.Context, postIds []int64) error {
	args := make([]interface{}, len(postIds))
	for i, id := range postIds {
		args[i] = id
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`delete from %s where id in (?`+strings.Repeat(",?", len(postIds)-1)+`)`, m.table)
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	return nil
}
