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
	wkfProcessFieldNames          = builder.RawFieldNames(&WkfProcess{})
	wkfProcessRows                = strings.Join(wkfProcessFieldNames, ",")
	wkfProcessRowsExpectAutoSet   = strings.Join(stringx.Remove(wkfProcessFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	wkfProcessRowsWithPlaceHolder = strings.Join(stringx.Remove(wkfProcessFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheWkfProcessIdPrefix = "cache:wkfProcess:id:"
)

type (
	wkfProcessModel interface {
		Insert(ctx context.Context, data *WkfProcess) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*WkfProcess, error)
		Update(ctx context.Context, data *WkfProcess) error
		Delete(ctx context.Context, id int64) error
	}

	defaultWkfProcessModel struct {
		sqlc.CachedConn
		table string
	}

	WkfProcess struct {
		Id        int64     `db:"id"`        // 编码
		Name      string    `db:"name"`      // 名称
		Icon      string    `db:"icon"`      // ICON
		Structure string    `db:"structure"` // 流程结构
		Classify  int64     `db:"classify"`  // 分类
		Template  string    `db:"template"`  // 模板
		Task      string    `db:"task"`      // 任务
		Notice    string    `db:"notice"`    // 通知
		Remark    string    `db:"remark"`    // 备注
		CreateBy  int64     `db:"create_by"` // 创建者
		UpdateBy  int64     `db:"update_by"` // 更新者
		CreateAt  time.Time `db:"create_at"` // 创建时间
		UpdateAt  time.Time `db:"update_at"` // 更新时间
	}
)

func newWkfProcessModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultWkfProcessModel {
	return &defaultWkfProcessModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`wkf_process`",
	}
}

func (m *defaultWkfProcessModel) withSession(session sqlx.Session) *defaultWkfProcessModel {
	return &defaultWkfProcessModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`wkf_process`",
	}
}

func (m *defaultWkfProcessModel) Delete(ctx context.Context, id int64) error {
	wkfProcessIdKey := fmt.Sprintf("%s%v", cacheWkfProcessIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, wkfProcessIdKey)
	return err
}

func (m *defaultWkfProcessModel) FindOne(ctx context.Context, id int64) (*WkfProcess, error) {
	wkfProcessIdKey := fmt.Sprintf("%s%v", cacheWkfProcessIdPrefix, id)
	var resp WkfProcess
	err := m.QueryRowCtx(ctx, &resp, wkfProcessIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wkfProcessRows, m.table)
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

func (m *defaultWkfProcessModel) Insert(ctx context.Context, data *WkfProcess) (sql.Result, error) {
	wkfProcessIdKey := fmt.Sprintf("%s%v", cacheWkfProcessIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, wkfProcessRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Icon, data.Structure, data.Classify, data.Template, data.Task, data.Notice, data.Remark, data.CreateBy, data.UpdateBy)
	}, wkfProcessIdKey)
	return ret, err
}

func (m *defaultWkfProcessModel) Update(ctx context.Context, data *WkfProcess) error {
	wkfProcessIdKey := fmt.Sprintf("%s%v", cacheWkfProcessIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, wkfProcessRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Icon, data.Structure, data.Classify, data.Template, data.Task, data.Notice, data.Remark, data.CreateBy, data.UpdateBy, data.Id)
	}, wkfProcessIdKey)
	return err
}

func (m *defaultWkfProcessModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheWkfProcessIdPrefix, primary)
}

func (m *defaultWkfProcessModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", wkfProcessRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultWkfProcessModel) tableName() string {
	return m.table
}