package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessRetrieveLogic {
	return &ProcessRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcessRetrieveLogic) ProcessRetrieve(in *wkf.ProcessRetrieveReq) (*wkf.ProcessRetrieveResp, error) {
	res, err := l.svcCtx.ProcessModel.FindOne(l.ctx, in.ProcessId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ProcessRetrieveResp{
		ProcessId: res.Id,
		Name:      res.Name,
		Icon:      res.Icon,
		Structure: res.Structure,
		Classify:  res.Classify,
		Template:  res.Template,
		Task:      res.Task,
		Notice:    res.Notice,
		Remark:    res.Remark,
	}, nil
}
