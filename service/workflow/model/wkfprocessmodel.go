package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfProcessModel = (*customWkfProcessModel)(nil)

type (
	// WkfProcessModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfProcessModel.
	WkfProcessModel interface {
		wkfProcessModel
		FindAll(context.Context, int64, int64) ([]*WkfProcess, error)
		Count(context.Context) (int64, error)
		FindProcessClassify(context.Context, string) ([]*WkfProcessClassify, error)
	}

	customWkfProcessModel struct {
		*defaultWkfProcessModel
		classifyTable string
	}

	WkfProcessClassify struct {
		Id           int64  `db:"id"`            // 编码
		Name         string `db:"name"`          // 名称
		Icon         string `db:"icon"`          // ICON
		ClassifyId   int64  `db:"classify_id"`   // 分类ID
		ClassifyName string `db:"classify_name"` // 分类名称
		Remark       string `db:"remark"`        // 备注
	}
)

// NewWkfProcessModel returns a model for the database table.
func NewWkfProcessModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfProcessModel {
	return &customWkfProcessModel{
		defaultWkfProcessModel: newWkfProcessModel(conn, c, opts...),
		classifyTable:          "`wkf_classify`",
	}
}

func (m *customWkfProcessModel) FindAll(ctx context.Context, Current int64, PageSize int64) ([]*WkfProcess, error) {
	var resp []*WkfProcess
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

func (m *customWkfProcessModel) Count(ctx context.Context) (int64, error) {
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

func (m *customWkfProcessModel) FindProcessClassify(ctx context.Context, name string) ([]*WkfProcessClassify, error) {
	var query string
	var resp []*WkfProcessClassify

	if name == "" {
		query = fmt.Sprintf("SELECT wp.id, wp.name, wp.icon, wp.remark, wc.id as classify_id, wc.name as classify_name from %s wp left join %s wc on wp.classify = wc.id ;", m.table, m.classifyTable)
	} else {
		query = fmt.Sprintf("SELECT wp.id, wp.name, wp.icon, wp.remark, wc.id as classify_id, wc.name as classify_name from %s wp left join %s wc on wp.classify = wc.id where wp.name like '%s';", m.table, m.classifyTable, name)
	}

	err := m.QueryRowsNoCacheCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
