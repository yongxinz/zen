package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/common/globalkey"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TaskListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TaskListLogic {
	return &TaskListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TaskListLogic) TaskList(in *wkf.TaskListReq) (*wkf.TaskListResp, error) {
	res, err := l.svcCtx.TaskModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get tasklist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.TaskModel.Count(l.ctx)

	var data []*wkf.TaskListData
	for _, item := range res {
		data = append(data, &wkf.TaskListData{
			TaskId:    item.Id,
			Name:      item.Name,
			Category:  item.Category,
			Content:   item.Content,
			Remark:    item.Remark,
			CreateBy:  item.CreateBy,
			UpdateBy:  item.UpdateBy,
			CreatedAt: item.CreateAt.Format(globalkey.SysDateFormat),
			UpdatedAt: item.UpdateAt.Format(globalkey.SysDateFormat),
		})
	}

	return &wkf.TaskListResp{
		Count: count,
		Data:  data,
	}, nil
}
