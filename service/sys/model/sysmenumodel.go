package model

import (
	"context"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysMenuModel = (*customSysMenuModel)(nil)

type (
	// SysMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysMenuModel.
	SysMenuModel interface {
		sysMenuModel
		FindMenuList(context.Context, []int64) ([]*SysMenu, error)
		FindAll(context.Context, int64, int64) ([]*SysMenu, error)
		Count(context.Context) (int64, error)
	}

	customSysMenuModel struct {
		*defaultSysMenuModel
	}
)

// NewSysMenuModel returns a model for the database table.
func NewSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysMenuModel {
	return &customSysMenuModel{
		defaultSysMenuModel: newSysMenuModel(conn, c, opts...),
	}
}

func (m *customSysMenuModel) FindMenuList(ctx context.Context, MenuIds []int64) ([]*SysMenu, error) {
	args := make([]interface{}, len(MenuIds))
	for i, id := range MenuIds {
		args[i] = id
	}

	var resp []*SysMenu
	query := fmt.Sprintf(`select %s from %s where id in (?`+strings.Repeat(",?", len(MenuIds)-1)+`)`, sysMenuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, args...)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysMenuModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysMenu, error) {
	var resp []*SysMenu
	query := fmt.Sprintf("select %s from %s order by id limit ?,?", sysMenuRows, m.table)
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

func (m *customSysMenuModel) Count(ctx context.Context) (int64, error) {
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
