package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketDeleteLogic {
	return &TicketDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketDeleteLogic) TicketDelete(in *wkf.TicketDeleteReq) (*wkf.TicketDeleteResp, error) {
	err := l.svcCtx.TicketModel.Delete(l.ctx, in.TicketId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TicketDeleteResp{}, nil
}
