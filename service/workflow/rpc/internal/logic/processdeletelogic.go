package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessDeleteLogic {
	return &ProcessDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcessDeleteLogic) ProcessDelete(in *wkf.ProcessDeleteReq) (*wkf.ProcessDeleteResp, error) {
	err := l.svcCtx.ProcessModel.Delete(l.ctx, in.ProcessId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ProcessDeleteResp{}, nil
}
