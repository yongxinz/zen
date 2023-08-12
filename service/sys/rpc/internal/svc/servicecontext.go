package svc

import (
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.SysUserModel
	MenuModel     model.SysMenuModel
	RoleModel     model.SysRoleModel
	RoleMenuModel model.SysRoleMenuModel
	DeptModel     model.SysDeptModel
	PostModel     model.SysPostModel
	LoginLogModel model.SysLoginLogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewSysUserModel(conn, c.CacheRedis),
		MenuModel:     model.NewSysMenuModel(conn, c.CacheRedis),
		RoleModel:     model.NewSysRoleModel(conn, c.CacheRedis),
		RoleMenuModel: model.NewSysRoleMenuModel(conn, c.CacheRedis),
		DeptModel:     model.NewSysDeptModel(conn, c.CacheRedis),
		PostModel:     model.NewSysPostModel(conn, c.CacheRedis),
		LoginLogModel: model.NewSysLoginLogModel(conn, c.CacheRedis),
	}
}
