package task

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskAddLogic {
	return &TaskAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskAddLogic) TaskAdd(req *types.TaskAddReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.CreateBy = userId
	req.UpdateBy = userId

	var task wkfclient.TaskAddReq
	err = copier.Copy(&task, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.WkfRpc.TaskAdd(l.ctx, &task)
	if err != nil {
		return err
	}

	return nil
}
