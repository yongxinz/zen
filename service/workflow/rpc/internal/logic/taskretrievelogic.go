package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskRetrieveLogic {
	return &TaskRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskRetrieveLogic) TaskRetrieve(in *wkf.TaskRetrieveReq) (*wkf.TaskRetrieveResp, error) {
	res, err := l.svcCtx.TaskModel.FindOne(l.ctx, in.TaskId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TaskRetrieveResp{
		TaskId:   res.Id,
		Name:     res.Name,
		Category: res.Category,
		Content:  res.Content,
		Remark:   res.Remark,
	}, nil
}
