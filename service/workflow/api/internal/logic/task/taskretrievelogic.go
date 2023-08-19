package task

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskRetrieveLogic {
	return &TaskRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskRetrieveLogic) TaskRetrieve(req *types.TaskRetrieveReq) (resp *types.TaskRetrieveResp, err error) {
	res, err := l.svcCtx.WkfRpc.TaskRetrieve(l.ctx, &wkfclient.TaskRetrieveReq{
		TaskId: req.TaskId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.TaskRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
