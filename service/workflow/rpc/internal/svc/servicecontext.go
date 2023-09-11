package svc

import (
	"github.com/yongxin/zen/service/sys/rpc/sysclient"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	ClassifyModel    model.WkfClassifyModel
	TemplateModel    model.WkfTemplateModel
	TaskModel        model.WkfTaskModel
	ProcessModel     model.WkfProcessModel
	CirculationModel model.WkfCirculationModel
	TicketModel      model.WkfTicketModel
	FormModel        model.WkfFormModel

	SysRpc sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,

		ClassifyModel:    model.NewWkfClassifyModel(conn, c.CacheRedis),
		TemplateModel:    model.NewWkfTemplateModel(conn, c.CacheRedis),
		TaskModel:        model.NewWkfTaskModel(conn, c.CacheRedis),
		ProcessModel:     model.NewWkfProcessModel(conn, c.CacheRedis),
		CirculationModel: model.NewWkfCirculationModel(conn, c.CacheRedis),
		TicketModel:      model.NewWkfTicketModel(conn, c.CacheRedis),
		FormModel:        model.NewWkfFormModel(conn, c.CacheRedis),

		SysRpc: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
