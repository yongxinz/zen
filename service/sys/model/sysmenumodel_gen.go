// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysMenuFieldNames          = builder.RawFieldNames(&SysMenu{})
	sysMenuRows                = strings.Join(sysMenuFieldNames, ",")
	sysMenuRowsExpectAutoSet   = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysMenuRowsWithPlaceHolder = strings.Join(stringx.Remove(sysMenuFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSysMenuIdPrefix = "cache:sysMenu:id:"
)

type (
	sysMenuModel interface {
		Insert(ctx context.Context, data *SysMenu) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysMenu, error)
		Update(ctx context.Context, data *SysMenu) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysMenuModel struct {
		sqlc.CachedConn
		table string
	}

	SysMenu struct {
		Id         int64          `db:"id"`
		MenuName   sql.NullString `db:"menu_name"`
		Title      sql.NullString `db:"title"`
		Icon       sql.NullString `db:"icon"`
		Path       sql.NullString `db:"path"`
		Paths      string         `db:"paths"`
		MenuType   sql.NullString `db:"menu_type"`
		Action     sql.NullString `db:"action"`
		Permission sql.NullString `db:"permission"`
		ParentId   sql.NullInt64  `db:"parent_id"`
		NoCache    sql.NullInt64  `db:"no_cache"`
		Breadcrumb sql.NullString `db:"breadcrumb"`
		Component  sql.NullString `db:"component"`
		Sort       sql.NullInt64  `db:"sort"`
		Visible    sql.NullString `db:"visible"`
		IsFrame    string         `db:"is_frame"`
		CreateBy   sql.NullInt64  `db:"create_by"`  // 创建者
		UpdateBy   sql.NullInt64  `db:"update_by"`  // 更新者
		CreatedAt  sql.NullTime   `db:"created_at"` // 创建时间
		UpdatedAt  sql.NullTime   `db:"updated_at"` // 最后更新时间
	}
)

func newSysMenuModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysMenuModel {
	return &defaultSysMenuModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_menu`",
	}
}

func (m *defaultSysMenuModel) withSession(session sqlx.Session) *defaultSysMenuModel {
	return &defaultSysMenuModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`sys_menu`",
	}
}

func (m *defaultSysMenuModel) Delete(ctx context.Context, id int64) error {
	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, sysMenuIdKey)
	return err
}

func (m *defaultSysMenuModel) FindOne(ctx context.Context, id int64) (*SysMenu, error) {
	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, id)
	var resp SysMenu
	err := m.QueryRowCtx(ctx, &resp, sysMenuIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysMenuModel) Insert(ctx context.Context, data *SysMenu) (sql.Result, error) {
	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, sysMenuRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.MenuName, data.Title, data.Icon, data.Path, data.Paths, data.MenuType, data.Action, data.Permission, data.ParentId, data.NoCache, data.Breadcrumb, data.Component, data.Sort, data.Visible, data.IsFrame, data.CreateBy, data.UpdateBy)
	}, sysMenuIdKey)
	return ret, err
}

func (m *defaultSysMenuModel) Update(ctx context.Context, data *SysMenu) error {
	sysMenuIdKey := fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysMenuRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.MenuName, data.Title, data.Icon, data.Path, data.Paths, data.MenuType, data.Action, data.Permission, data.ParentId, data.NoCache, data.Breadcrumb, data.Component, data.Sort, data.Visible, data.IsFrame, data.CreateBy, data.UpdateBy, data.Id)
	}, sysMenuIdKey)
	return err
}

func (m *defaultSysMenuModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSysMenuIdPrefix, primary)
}

func (m *defaultSysMenuModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysMenuRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysMenuModel) tableName() string {
	return m.table
}