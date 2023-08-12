package svc

import (
	"github.com/yongxin/zen/service/workflow/api/internal/config"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	WkfRpc wkfclient.Wkf
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		WkfRpc: wkfclient.NewWkf(zrpc.MustNewClient(c.WkfRpc)),
	}
}
