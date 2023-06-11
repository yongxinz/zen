package model

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleMenuModel = (*customSysRoleMenuModel)(nil)

type (
	// SysRoleMenuModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleMenuModel.
	SysRoleMenuModel interface {
		sysRoleMenuModel
		FindMenuIds(context.Context, int64) ([]*SysRoleMenu, error)
		DeleteByRoleId(context.Context, int64) error
	}

	customSysRoleMenuModel struct {
		*defaultSysRoleMenuModel
	}
)

// NewSysRoleMenuModel returns a model for the database table.
func NewSysRoleMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) SysRoleMenuModel {
	return &customSysRoleMenuModel{
		defaultSysRoleMenuModel: newSysRoleMenuModel(conn, c, opts...),
	}
}

func (m *customSysRoleMenuModel) FindMenuIds(ctx context.Context, RoleId int64) ([]*SysRoleMenu, error) {
	var resp []*SysRoleMenu
	query := fmt.Sprintf("select %s from %s where role_id = ?", sysRoleMenuRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, RoleId)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customSysRoleMenuModel) DeleteByRoleId(ctx context.Context, roleId int64) error {
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `role_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, roleId)
	})
	if err != nil {
		return err
	}
	return nil
}
