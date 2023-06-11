package menu

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList() (resp []*types.MenuListResp, err error) {
	res, err := l.svcCtx.SysRpc.MenuList(l.ctx, &sysclient.MenuListReq{})
	if err != nil {
		return nil, err
	}

	var data []*types.MenuListResp
	copier.Copy(&data, res.Data)
	resp = data

	return
}
