package role

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRetrieveLogic {
	return &RoleRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleRetrieveLogic) RoleRetrieve(req *types.RoleRetrieveReq) (resp *types.RoleRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.RoleRetrieve(l.ctx, &sysclient.RoleRetrieveReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.RoleRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
