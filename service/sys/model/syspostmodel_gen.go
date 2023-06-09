// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysPostFieldNames          = builder.RawFieldNames(&SysPost{})
	sysPostRows                = strings.Join(sysPostFieldNames, ",")
	sysPostRowsExpectAutoSet   = strings.Join(stringx.Remove(sysPostFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	sysPostRowsWithPlaceHolder = strings.Join(stringx.Remove(sysPostFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheSysPostIdPrefix = "cache:sysPost:id:"
)

type (
	sysPostModel interface {
		Insert(ctx context.Context, data *SysPost) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysPost, error)
		Update(ctx context.Context, data *SysPost) error
		Delete(ctx context.Context, id int64) error
	}

	defaultSysPostModel struct {
		sqlc.CachedConn
		table string
	}

	SysPost struct {
		Id        int64     `db:"id"`         // 编码
		PostName  string    `db:"post_name"`  // 岗位名称
		PostCode  string    `db:"post_code"`  // 岗位编码
		Sort      int64     `db:"sort"`       // 排序
		Remark    string    `db:"remark"`     // 备注
		Status    int64     `db:"status"`     // 状态
		CreateBy  int64     `db:"create_by"`  // 创建者
		UpdateBy  int64     `db:"update_by"`  // 更新者
		CreatedAt time.Time `db:"created_at"` // 创建时间
		UpdatedAt time.Time `db:"updated_at"` // 更新时间
	}
)

func newSysPostModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultSysPostModel {
	return &defaultSysPostModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`sys_post`",
	}
}

func (m *defaultSysPostModel) withSession(session sqlx.Session) *defaultSysPostModel {
	return &defaultSysPostModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`sys_post`",
	}
}

func (m *defaultSysPostModel) Delete(ctx context.Context, id int64) error {
	sysPostIdKey := fmt.Sprintf("%s%v", cacheSysPostIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, sysPostIdKey)
	return err
}

func (m *defaultSysPostModel) FindOne(ctx context.Context, id int64) (*SysPost, error) {
	sysPostIdKey := fmt.Sprintf("%s%v", cacheSysPostIdPrefix, id)
	var resp SysPost
	err := m.QueryRowCtx(ctx, &resp, sysPostIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysPostRows, m.table)
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

func (m *defaultSysPostModel) Insert(ctx context.Context, data *SysPost) (sql.Result, error) {
	sysPostIdKey := fmt.Sprintf("%s%v", cacheSysPostIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, sysPostRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.PostName, data.PostCode, data.Sort, data.Remark, data.Status, data.CreateBy, data.UpdateBy)
	}, sysPostIdKey)
	return ret, err
}

func (m *defaultSysPostModel) Update(ctx context.Context, data *SysPost) error {
	sysPostIdKey := fmt.Sprintf("%s%v", cacheSysPostIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, sysPostRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.PostName, data.PostCode, data.Sort, data.Remark, data.Status, data.CreateBy, data.UpdateBy, data.Id)
	}, sysPostIdKey)
	return err
}

func (m *defaultSysPostModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheSysPostIdPrefix, primary)
}

func (m *defaultSysPostModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", sysPostRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysPostModel) tableName() string {
	return m.table
}
