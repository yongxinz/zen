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

type ProcessListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessListLogic {
	return &ProcessListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcessListLogic) ProcessList(in *wkf.ProcessListReq) (*wkf.ProcessListResp, error) {
	res, err := l.svcCtx.ProcessModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get processlist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.ProcessModel.Count(l.ctx)

	var data []*wkf.ProcessListData
	for _, item := range res {
		data = append(data, &wkf.ProcessListData{
			ProcessId: item.Id,
			Name:      item.Name,
			Icon:      item.Icon,
			Structure: item.Structure,
			Classify:  item.Classify,
			Template:  item.Template,
			Task:      item.Task,
			Notice:    item.Notice,
			Remark:    item.Remark,
			CreateBy:  item.CreateBy,
			UpdateBy:  item.UpdateBy,
			CreatedAt: item.CreateAt.Format(globalkey.SysDateFormat),
			UpdatedAt: item.UpdateAt.Format(globalkey.SysDateFormat),
		})
	}

	return &wkf.ProcessListResp{
		Count: count,
		Data:  data,
	}, nil
}
