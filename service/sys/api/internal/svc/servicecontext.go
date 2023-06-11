package svc

import (
	"github.com/yongxin/zen/service/sys/api/internal/config"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	SysRpc sysclient.Sys
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		SysRpc: sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
	}
}
