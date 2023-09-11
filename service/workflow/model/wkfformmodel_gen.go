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
	wkfFormFieldNames          = builder.RawFieldNames(&WkfForm{})
	wkfFormRows                = strings.Join(wkfFormFieldNames, ",")
	wkfFormRowsExpectAutoSet   = strings.Join(stringx.Remove(wkfFormFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	wkfFormRowsWithPlaceHolder = strings.Join(stringx.Remove(wkfFormFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheWkfFormIdPrefix = "cache:wkfForm:id:"
)

type (
	wkfFormModel interface {
		Insert(ctx context.Context, data *WkfForm) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*WkfForm, error)
		Update(ctx context.Context, data *WkfForm) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWkfFormModel struct {
		sqlc.CachedConn
		table string
	}

	WkfForm struct {
		Id            int64     `db:"id"`             // 编码
		TicketId      int64     `db:"ticket_id"`      // 工单ID
		FormStructure string    `db:"form_structure"` // 表单结构
		FormData      string    `db:"form_data"`      // 表单数据
		CreateAt      time.Time `db:"create_at"`      // 创建时间
		UpdateAt      time.Time `db:"update_at"`      // 更新时间
	}
)

func newWkfFormModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultWkfFormModel {
	return &defaultWkfFormModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`wkf_form`",
	}
}

func (m *defaultWkfFormModel) withSession(session sqlx.Session) *defaultWkfFormModel {
	return &defaultWkfFormModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`wkf_form`",
	}
}

func (m *defaultWkfFormModel) Delete(ctx context.Context, id int64) error {
	wkfFormIdKey := fmt.Sprintf("%s%v", cacheWkfFormIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, wkfFormIdKey)
	return err
}

func (m *defaultWkfFormModel) FindOne(ctx context.Context, id int64) (*WkfForm, error) {
	wkfFormIdKey := fmt.Sprintf("%s%v", cacheWkfFormIdPrefix, id)
	var resp WkfForm
	err := m.QueryRowCtx(ctx, &resp, wkfFormIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wkfFormRows, m.table)
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

func (m *defaultWkfFormModel) Insert(ctx context.Context, data *WkfForm) (sql.Result, error) {
	wkfFormIdKey := fmt.Sprintf("%s%v", cacheWkfFormIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, wkfFormRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.TicketId, data.FormStructure, data.FormData)
	}, wkfFormIdKey)
	return ret, err
}

func (m *defaultWkfFormModel) Update(ctx context.Context, data *WkfForm) error {
	wkfFormIdKey := fmt.Sprintf("%s%v", cacheWkfFormIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wkfFormRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.TicketId, data.FormStructure, data.FormData, data.Id)
	}, wkfFormIdKey)
	return err
}

func (m *defaultWkfFormModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheWkfFormIdPrefix, primary)
}

func (m *defaultWkfFormModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wkfFormRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultWkfFormModel) tableName() string {
	return m.table
}