package logic

import (
	"context"
	"fmt"

	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketFinishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketFinishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketFinishLogic {
	return &TicketFinishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketFinishLogic) TicketFinish(in *wkf.TicketFinishReq) (*wkf.TicketFinishResp, error) {
	// get ticket info
	ticketInfo, err := l.svcCtx.TicketModel.FindOne(l.ctx, in.TicketId)
	if err != nil {
		err = fmt.Errorf("TicketFinish Ticket FindOne, %v", err.Error())
		return &wkf.TicketFinishResp{}, err
	}

	if ticketInfo.IsEnd == 1 {
		err = fmt.Errorf("ticket already finished, %v", err.Error())
		return &wkf.TicketFinishResp{}, err
	}

	ticketInfo.IsEnd = 1
	err = l.svcCtx.TicketModel.Update(l.ctx, ticketInfo)
	if err != nil {
		err = fmt.Errorf("TicketFinish Ticket Update error, %v", err.Error())
		return &wkf.TicketFinishResp{}, err
	}

	wkfCirculation := &model.WkfCirculation{
		TicketId:    in.TicketId,
		State:       "结束工单",
		Circulation: "结束",
		HandlerId:   in.UpdateBy,
		HandlerName: "xxx",
		Remark:      "手动结束工单",
		Status:      2,
	}
	_, err = l.svcCtx.CirculationModel.Insert(l.ctx, wkfCirculation)
	if err != nil {
		return &wkf.TicketFinishResp{}, err
	}

	return &wkf.TicketFinishResp{}, nil
}
