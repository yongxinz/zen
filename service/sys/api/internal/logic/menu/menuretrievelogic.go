package menu

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuRetrieveLogic {
	return &MenuRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuRetrieveLogic) MenuRetrieve(req *types.MenuRetrieveReq) (resp *types.MenuRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.MenuRetrieve(l.ctx, &sysclient.MenuRetrieveReq{
		MenuId: req.MenuId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.MenuRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
