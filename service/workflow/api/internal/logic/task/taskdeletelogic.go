package task

import (
	"context"

	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskDeleteLogic {
	return &TaskDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskDeleteLogic) TaskDelete(req *types.TaskDeleteReq) error {
	_, err := l.svcCtx.WkfRpc.TaskDelete(l.ctx, &wkfclient.TaskDeleteReq{
		TaskId: req.TaskId,
	})

	return err
}
