package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskUpdateLogic {
	return &TaskUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskUpdateLogic) TaskUpdate(in *wkf.TaskUpdateReq) (*wkf.TaskUpdateResp, error) {
	task, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(task, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.TaskModel.Update(l.ctx, task)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TaskUpdateResp{}, nil
}
