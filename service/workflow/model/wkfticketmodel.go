package model

import (
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WkfTicketModel = (*customWkfTicketModel)(nil)

type (
	// WkfTicketModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWkfTicketModel.
	WkfTicketModel interface {
		wkfTicketModel
		FindAll(context.Context, int64, int64, map[string]int64) ([]*TicketList, error)
		Count(context.Context, map[string]int64) (int64, error)
	}

	customWkfTicketModel struct {
		*defaultWkfTicketModel
		processTable string
	}

	TicketList struct {
		WkfTicket
		ProcessName string `db:"process_name"`
	}
)

// NewWkfTicketModel returns a model for the database table.
func NewWkfTicketModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WkfTicketModel {
	return &customWkfTicketModel{
		defaultWkfTicketModel: newWkfTicketModel(conn, c, opts...),
		processTable:          "wkf_process",
	}
}

func (m *customWkfTicketModel) FindAll(ctx context.Context, Current, PageSize int64, params map[string]int64) ([]*TicketList, error) {
	var (
		resp []*TicketList
		err  error
	)
	category := params["category"]

	switch category {
	case 1:
		query := fmt.Sprintf(
			`SELECT
				wt.*, wp.name as process_name
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			where
				(JSON_CONTAINS(wt.state, JSON_OBJECT('processor', ?))
				and JSON_CONTAINS(wt.state, JSON_OBJECT('process_method', 'person')))
			or (JSON_CONTAINS(wt.state, JSON_OBJECT('processor', ?))
				and JSON_CONTAINS(wt.state, JSON_OBJECT('process_method', 'role')))
			or (JSON_CONTAINS(wt.state, JSON_OBJECT('processor', ?))
				and JSON_CONTAINS(wt.state, JSON_OBJECT('process_method', 'department')))
			limit ?, ?;`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, params["userId"], params["roleId"], params["deptId"],
			(Current-1)*PageSize, PageSize)
	case 2:
		query := fmt.Sprintf(
			`SELECT
				wt.*, wp.name as process_name
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			where
				wt.create_by = ?
			limit ?, ?;`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, params["userId"], (Current-1)*PageSize, PageSize)
	case 3:
		query := fmt.Sprintf(
			`SELECT
				wt.*, wp.name as process_name
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			where
				JSON_CONTAINS(wt.related_person, JSON_ARRAY(?))
			limit ?, ?;`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, params["userId"], (Current-1)*PageSize, PageSize)
	case 4:
		query := fmt.Sprintf(
			`SELECT
				wt.*, wp.name as process_name
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			limit ?, ?;`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, (Current-1)*PageSize, PageSize)
	default:
		return nil, fmt.Errorf("请确认查询的数据类型是否正确")
	}

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customWkfTicketModel) Count(ctx context.Context, params map[string]int64) (int64, error) {
	var (
		count int64
		err   error
	)
	category := params["category"]

	switch category {
	case 1:
		query := fmt.Sprintf(
			`SELECT
				count(*)
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			where
				(JSON_CONTAINS(wt.state, JSON_OBJECT('processor', ?))
				and JSON_CONTAINS(wt.state, JSON_OBJECT('process_method', 'person')))
			or (JSON_CONTAINS(wt.state, JSON_OBJECT('processor', ?))
				and JSON_CONTAINS(wt.state, JSON_OBJECT('process_method', 'role')))
			or (JSON_CONTAINS(wt.state, JSON_OBJECT('processor', ?))
				and JSON_CONTAINS(wt.state, JSON_OBJECT('process_method', 'department')));`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &count, query, params["userId"], params["roleId"], params["deptId"])
	case 2:
		query := fmt.Sprintf(
			`SELECT
				count(*)
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			where
				wt.create_by = ?;`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &count, query, params["userId"])
	case 3:
		query := fmt.Sprintf(
			`SELECT
				count(*)
			from
				%s wt
			left join %s wp on
				wt.process_id = wp.id
			where
				JSON_CONTAINS(wt.related_person, JSON_ARRAY(?));`, m.table, m.processTable)
		err = m.QueryRowsNoCacheCtx(ctx, &count, query, params["userId"])
	case 4:
	default:
		return 0, fmt.Errorf("请确认查询的数据类型是否正确")
	}

	switch err {
	case nil:
		return count, nil
	default:
		return count, err
	}
}
