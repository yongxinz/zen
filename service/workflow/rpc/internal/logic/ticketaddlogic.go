package logic

import (
	"context"
	"encoding/json"
	"strconv"

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
		processState ProcessState
		nodeState    []interface{}
		template     map[string][]interface{}
	)

	processInfo, err := l.svcCtx.ProcessModel.FindOne(l.ctx, in.ProcessId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	err = json.Unmarshal([]byte(processInfo.Structure), &processState.Structure)
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
		Target:      nodeState[0].(map[string]interface{})["id"].(string),
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

type ProcessState struct {
	Structure map[string][]map[string]interface{}
}

// 获取节点信息
func (p *ProcessState) GetNode(stateId string) (nodeValue map[string]interface{}, err error) {
	for _, node := range p.Structure["nodes"] {
		if node["id"] == stateId {
			nodeValue = node
		}
	}
	return
}

// 获取流转信息
func (p *ProcessState) GetEdge(stateId string, classify string) (edgeValue []map[string]interface{}, err error) {
	var (
		leftSort  int
		rightSort int
	)

	for _, edge := range p.Structure["edges"] {
		if edge[classify] == stateId {
			edgeValue = append(edgeValue, edge)
		}
	}

	// 排序
	if len(edgeValue) > 1 {
		for i := 0; i < len(edgeValue)-1; i++ {
			for j := i + 1; j < len(edgeValue); j++ {
				if t, ok := edgeValue[i]["sort"]; ok {
					leftSort, _ = strconv.Atoi(t.(string))
				}
				if t, ok := edgeValue[j]["sort"]; ok {
					rightSort, _ = strconv.Atoi(t.(string))
				}
				if leftSort > rightSort {
					edgeValue[j], edgeValue[i] = edgeValue[i], edgeValue[j]
				}
			}
		}
	}

	return
}
