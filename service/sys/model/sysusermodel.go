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

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		FindAll(context.Context, int64, int64) ([]*SysUserList, error)
		Count(context.Context) (int64, error)
		DeleteMulti(context.Context, []int64) error
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}

	SysUserList struct {
		SysUser
		SysDept
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn, c cache.CacheConf) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn, c),
	}
}

func (m *customSysUserModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*SysUserList, error) {
	var resp []*SysUserList
	query := "select sys_user.*, sd.* from sys_user left join sys_dept sd on sys_user.dept_id = sd.id limit ?,?"
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

func (m *customSysUserModel) Count(ctx context.Context) (int64, error) {
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

func (m *customSysUserModel) DeleteMulti(ctx context.Context, userIds []int64) error {
	args := make([]interface{}, len(userIds))
	for i, id := range userIds {
		args[i] = id
	}

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf(`delete from %s where id in (?`+strings.Repeat(",?", len(userIds)-1)+`)`, m.table)
		return conn.ExecCtx(ctx, query, args...)
	})
	if err != nil {
		return err
	}
	return nil
}
