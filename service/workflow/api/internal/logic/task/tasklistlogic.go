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

type TaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TaskListLogic) TaskList(req *types.TaskListReq) (resp *types.TaskListResp, err error) {
	res, err := l.svcCtx.WkfRpc.TaskList(l.ctx, &wkfclient.TaskListReq{
		Page: req.PageIndex,
		Size: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get tasklist error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var task types.TaskListData
	var list []types.TaskListData
	for _, v := range res.Data {
		copier.Copy(&task, &v)
		list = append(list, task)
	}

	resp = &types.TaskListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
