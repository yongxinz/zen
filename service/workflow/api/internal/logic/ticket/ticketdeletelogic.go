package ticket

import (
	"context"

	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTicketDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketDeleteLogic {
	return &TicketDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TicketDeleteLogic) TicketDelete(req *types.TicketDeleteReq) error {
	_, err := l.svcCtx.WkfRpc.TicketDelete(l.ctx, &wkfclient.TicketDeleteReq{
		TicketId: req.TicketId,
	})

	return err
}
