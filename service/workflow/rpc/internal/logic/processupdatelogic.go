package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessUpdateLogic {
	return &ProcessUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcessUpdateLogic) ProcessUpdate(in *wkf.ProcessUpdateReq) (*wkf.ProcessUpdateResp, error) {
	process, err := l.svcCtx.ProcessModel.FindOne(l.ctx, in.ProcessId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(process, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.ProcessModel.Update(l.ctx, process)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ProcessUpdateResp{}, nil
}
