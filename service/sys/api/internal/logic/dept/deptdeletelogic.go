package dept

import (
	"context"

	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeptDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptDeleteLogic {
	return &DeptDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeptDeleteLogic) DeptDelete(req *types.DeptDeleteReq) error {
	_, err := l.svcCtx.SysRpc.DeptDelete(l.ctx, &sysclient.DeptDeleteReq{
		DeptId: req.DeptId,
	})
	return err
}
