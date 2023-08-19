package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskAddLogic {
	return &TaskAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskAddLogic) TaskAdd(in *wkf.TaskAddReq) (*wkf.TaskAddResp, error) {
	var wkfTask = new(model.WkfTask)
	err := copier.Copy(wkfTask, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.TaskModel.Insert(l.ctx, wkfTask)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TaskAddResp{}, nil
}
