package logic

import (
	"context"
	"fmt"
	"time"

	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketUrgeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketUrgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketUrgeLogic {
	return &TicketUrgeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketUrgeLogic) TicketUrge(in *wkf.TicketUrgeReq) (*wkf.TicketUrgeResp, error) {
	// get ticket info
	ticketInfo, err := l.svcCtx.TicketModel.FindOne(l.ctx, in.TicketId)
	if err != nil {
		err = fmt.Errorf("TicketUrge Ticket FindOne, %v", err.Error())
		return &wkf.TicketUrgeResp{}, err
	}

	if ticketInfo.UrgeLasttime != 0 && (int64(time.Now().Unix())-ticketInfo.UrgeLasttime) < 600 {
		err = fmt.Errorf("can only be done once every ten minutes")
		return &wkf.TicketUrgeResp{}, err
	}

	// var state map[string]interface{}
	// err = json.Unmarshal([]byte(ticketInfo.State), &state)
	// if err != nil {
	// 	err = fmt.Errorf("TicketUrge Unmarshal to state error, %v", err.Error())
	// 	return nil, err
	// }

	ticketInfo.UrgeCount = ticketInfo.UrgeCount + 1
	ticketInfo.UrgeLasttime = int64(time.Now().Unix())

	err = l.svcCtx.TicketModel.Update(l.ctx, ticketInfo)
	if err != nil {
		err = fmt.Errorf("TicketTransfer Ticket Update error, %v", err.Error())
		return &wkf.TicketUrgeResp{}, err
	}

	return &wkf.TicketUrgeResp{}, nil
}
