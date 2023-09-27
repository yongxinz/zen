package logic

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketTransferLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketTransferLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketTransferLogic {
	return &TicketTransferLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketTransferLogic) TicketTransfer(in *wkf.TicketTransferReq) (*wkf.TicketTransferResp, error) {
	// get ticket info
	ticketInfo, err := l.svcCtx.TicketModel.FindOne(l.ctx, in.TicketId)
	if err != nil {
		err = fmt.Errorf("TicketTransfer Ticket FindOne, %v", err.Error())
		return &wkf.TicketTransferResp{}, err
	}

	var state []map[string]interface{}
	err = json.Unmarshal([]byte(ticketInfo.State), &state)
	if err != nil {
		err = fmt.Errorf("TicketTransfer Unmarshal to state error, %v", err.Error())
		return nil, err
	}

	state[0]["processor"] = []interface{}{in.UserId}
	state[0]["process_method"] = "person"

	stateValue, err := json.Marshal(state)
	if err != nil {
		err = fmt.Errorf("TicketTransfer Marshal to stateValue error, %v", err.Error())
		return nil, err
	}

	ticketInfo.State = string(stateValue)
	ticketInfo.UpdateBy = in.UpdateBy
	err = l.svcCtx.TicketModel.Update(l.ctx, ticketInfo)
	if err != nil {
		err = fmt.Errorf("TicketTransfer Ticket Update error, %v", err.Error())
		return &wkf.TicketTransferResp{}, err
	}

	wkfCirculation := &model.WkfCirculation{
		TicketId:    in.TicketId,
		State:       state[0]["label"].(string),
		Circulation: "转交",
		HandlerId:   in.UpdateBy,
		HandlerName: "xxx",
		Remark:      "",
		Status:      2,
	}
	_, err = l.svcCtx.CirculationModel.Insert(l.ctx, wkfCirculation)
	if err != nil {
		return &wkf.TicketTransferResp{}, err
	}

	return &wkf.TicketTransferResp{}, nil
}
