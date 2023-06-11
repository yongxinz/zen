package dept

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptTreeLogic {
	return &DeptTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptTreeLogic) DeptTree() (resp []*types.DeptTreeResp, err error) {
	res, err := l.svcCtx.SysRpc.DeptTree(l.ctx, &sysclient.DeptTreeReq{})
	if err != nil {
		return nil, err
	}

	var data []*types.DeptTreeResp
	copier.Copy(&data, res.Data)
	resp = data

	return
}
