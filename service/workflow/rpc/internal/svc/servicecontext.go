package svc

import (
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	ClassifyModel model.WkfClassifyModel
	TemplateModel model.WkfTemplateModel
	TaskModel     model.WkfTaskModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		ClassifyModel: model.NewWkfClassifyModel(conn, c.CacheRedis),
		TemplateModel: model.NewWkfTemplateModel(conn, c.CacheRedis),
		TaskModel:     model.NewWkfTaskModel(conn, c.CacheRedis),
	}
}
