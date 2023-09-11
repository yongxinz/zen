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

type TicketListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTicketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketListLogic {
	return &TicketListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TicketListLogic) TicketList(req *types.TicketListReq) (resp *types.TicketListResp, err error) {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	res, err := l.svcCtx.WkfRpc.TicketList(l.ctx, &wkfclient.TicketListReq{
		Page:     req.PageIndex,
		Size:     req.PageSize,
		Category: req.Category,
		UserId:   userId,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get ticketlist error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var data types.TicketListData
	var list []types.TicketListData
	for _, v := range res.Data {
		copier.Copy(&data, &v)
		list = append(list, data)
	}

	resp = &types.TicketListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
