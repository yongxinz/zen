package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/sys"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	targetState map[string]interface{}
	updateData  map[string]interface{}
	ProcessStructure
	UserPermission
}

func NewTicketHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketHandleLogic {
	return &TicketHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketHandleLogic) TicketHandle(in *wkf.TicketHandleReq) (*wkf.TicketHandleResp, error) {
	var err error

	// get ticket info
	l.UserPermission.TicketInfo, err = l.svcCtx.TicketModel.FindOne(l.ctx, in.TicketId)
	if err != nil {
		err = fmt.Errorf("TicketHandle Ticket FindOne, %v", err.Error())
		return &wkf.TicketHandleResp{}, err
	}

	// get user info
	l.UserPermission.UserInfo, err = l.svcCtx.SysRpc.UserRetrieve(l.ctx, &sys.UserInfoReq{
		UserId: in.UpdateBy,
	})
	if err != nil {
		err = fmt.Errorf("TicketHandle SysRpc UserRetrieve, %v", err.Error())
		return &wkf.TicketHandleResp{}, err
	}

	// Determine whether the user has operating permissions for the ticket
	userPermission, err := l.UserPermission.JudgeUserPermission()
	if err != nil || !userPermission {
		return nil, err
	}

	// get source state
	sourceState, err := l.getSourceState()
	if err != nil {
		return nil, err
	}

	// search process info
	processInfo, err := l.svcCtx.ProcessModel.FindOne(l.ctx, l.UserPermission.TicketInfo.ProcessId)
	if err != nil {
		err = fmt.Errorf("TicketHandle Process FindOne, %v", err.Error())
		return &wkf.TicketHandleResp{}, err
	}

	err = json.Unmarshal([]byte(processInfo.Structure), &l.ProcessStructure.Structure)
	if err != nil {
		err = fmt.Errorf("TicketHandle Unmarshal to processStructure error, %v", err.Error())
		return nil, err
	}
	var targetId, circulationValue string
	edges := l.ProcessStructure.Structure["edges"]
	for _, edge := range edges {
		flowProperties, _ := strconv.Atoi(edge["flowProperties"].(string))
		if edge["source"].(string) == sourceState["id"] && in.FlowProperties == int64(flowProperties) {
			targetId = edge["target"].(string)
			circulationValue = edge["label"].(string)
			break
		}
	}

	l.targetState, err = l.ProcessStructure.GetNode(targetId)
	if err != nil {
		return nil, err
	}

	stateValue := map[string]interface{}{
		"label": l.targetState["label"].(string),
		"id":    l.targetState["id"].(string),
	}
	l.updateData = map[string]interface{}{}

	switch l.targetState["clazz"].(string) {
	case "userTask":
		stateValue["processor"] = l.targetState["assignValue"].([]interface{})
		stateValue["process_method"] = l.targetState["assignType"].(string)
		l.updateData["state"] = stateValue
		err = l.hanlde(in.TicketId, in.FlowProperties)
		if err != nil {
			return nil, err
		}
	case "end":
		stateValue["processor"] = []int{}
		stateValue["process_method"] = ""
		l.updateData["state"] = stateValue
		err = l.hanlde(in.TicketId, in.FlowProperties)
		if err != nil {
			return nil, err
		}
	}

	wkfCirculation := &model.WkfCirculation{
		TicketId:    in.TicketId,
		State:       stateValue["label"].(string),
		Source:      stateValue["id"].(string),
		Target:      l.targetState["id"].(string),
		Circulation: circulationValue,
		HandlerId:   in.UpdateBy,
		HandlerName: "xxx",
		Status:      in.FlowProperties,
	}
	_, err = l.svcCtx.CirculationModel.Insert(l.ctx, wkfCirculation)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TicketHandleResp{}, err
}

func (l *TicketHandleLogic) hanlde(ticketId, flowProperties int64) error {
	if flowProperties == 0 {
		err := l.circulation(ticketId, flowProperties)
		return err
	}
	err := l.circulation(ticketId, flowProperties)
	return err
}

func (l *TicketHandleLogic) circulation(ticketId, flowProperties int64) error {
	state, err := json.Marshal(l.updateData["state"])
	if err != nil {
		err = fmt.Errorf("circulation Marshal state error, %v", err.Error())
		return err
	}

	ticket, err := l.svcCtx.TicketModel.FindOne(l.ctx, ticketId)
	if err != nil {
		err = fmt.Errorf("circulation TicketModel FindOne error, %v", err.Error())
		return err
	}
	ticket.State = string(state)
	ticket.IsDenied = flowProperties
	ticket.RelatedPerson = "xxxx"

	if l.targetState["clazz"].(string) == "end" {
		ticket.IsEnd = 1
	}

	err = l.svcCtx.TicketModel.Update(l.ctx, ticket)
	if err != nil {
		err = fmt.Errorf("circulation TicketModel Update error, %v", err.Error())
		return err
	}

	return nil
}

func (l *TicketHandleLogic) getSourceState() (map[string]interface{}, error) {
	var state []map[string]interface{}

	err := json.Unmarshal([]byte(l.UserPermission.TicketInfo.State), &state)
	if err != nil {
		err = fmt.Errorf("getSourceState Unmarshal to state error, %v", err.Error())
		return nil, err
	}

	return state[0], nil
}
