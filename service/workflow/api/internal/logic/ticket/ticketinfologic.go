package ticket

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTicketInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketInfoLogic {
	return &TicketInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TicketInfoLogic) TicketInfo(req *types.TicketInfoReq) (resp *types.TicketInfoResp, err error) {
	res, err := l.svcCtx.WkfRpc.TicketProcess(l.ctx, &wkfclient.TicketProcessReq{
		ProcessId: req.ProcessId,
		TicketId:  req.TicketId,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get ticketinfo error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	resp = &types.TicketInfoResp{}
	copier.Copy(resp, res)
	fmt.Println(resp)

	return
}
