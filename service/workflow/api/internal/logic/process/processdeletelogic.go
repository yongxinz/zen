package process

import (
	"context"

	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessDeleteLogic {
	return &ProcessDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessDeleteLogic) ProcessDelete(req *types.ProcessDeleteReq) error {
	_, err := l.svcCtx.WkfRpc.ProcessDelete(l.ctx, &wkfclient.ProcessDeleteReq{
		ProcessId: req.ProcessId,
	})

	return err
}
