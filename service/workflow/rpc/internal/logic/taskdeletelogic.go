package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskDeleteLogic {
	return &TaskDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskDeleteLogic) TaskDelete(in *wkf.TaskDeleteReq) (*wkf.TaskDeleteResp, error) {
	err := l.svcCtx.TaskModel.Delete(l.ctx, in.TaskId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TaskDeleteResp{}, nil
}
