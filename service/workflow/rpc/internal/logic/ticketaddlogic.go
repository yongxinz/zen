package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger

	ProcessStructure
}

func NewTicketAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketAddLogic {
	return &TicketAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketAddLogic) TicketAdd(in *wkf.TicketAddReq) (*wkf.TicketAddResp, error) {
	var (
		nodeState map[string]interface{}
		template  map[string][]interface{}
	)

	processInfo, err := l.svcCtx.ProcessModel.FindOne(l.ctx, in.ProcessId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	err = json.Unmarshal([]byte(processInfo.Structure), &l.ProcessStructure.Structure)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = json.Unmarshal([]byte(in.State), &nodeState)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	// nodeValue, err := processState.GetNode(nodeState[0].(map[string]interface{})["id"].(string))
	// if err != nil {
	// 	return
	// }

	wkfTicket := &model.WkfTicket{
		ProcessId:     in.ProcessId,
		ClassifyId:    in.ClassifyId,
		State:         in.State,
		RelatedPerson: "xxx",
		CreateBy:      in.CreateBy,
		UpdateBy:      in.UpdateBy,
	}
	res, err := l.svcCtx.TicketModel.Insert(l.ctx, wkfTicket)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	ticketId, _ := res.LastInsertId()

	err = json.Unmarshal([]byte(in.Template), &template)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	for i := 0; i < len(template["form_structure"]); i++ {
		formDataByte, _ := json.Marshal(template["form_data"][0])
		formStructureByte, _ := json.Marshal(template["form_structure"][0])
		wkfForm := &model.WkfForm{
			TicketId:      ticketId,
			FormData:      string(formDataByte),
			FormStructure: string(formStructureByte),
		}
		_, err := l.svcCtx.FormModel.Insert(l.ctx, wkfForm)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
	}

	wkfCirculation := &model.WkfCirculation{
		TicketId:    ticketId,
		State:       in.SourceState,
		Source:      in.Source,
		Target:      nodeState["id"].(string),
		Circulation: "新建",
		HandlerId:   in.CreateBy,
		HandlerName: "xxx",
		Status:      2,
	}
	_, err = l.svcCtx.CirculationModel.Insert(l.ctx, wkfCirculation)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TicketAddResp{}, nil
}
