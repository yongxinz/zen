package ticket

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

type TicketFinishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTicketFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketFinishLogic {
	return &TicketFinishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TicketFinishLogic) TicketFinish(req *types.TicketFinishReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.UpdateBy = userId

	var data wkfclient.TicketFinishReq
	err = copier.Copy(&data, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.WkfRpc.TicketFinish(l.ctx, &data)
	if err != nil {
		return err
	}

	return nil
}
