package logic

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"

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
	FormData [][]byte
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
		nodeState []interface{}
		template  map[string][]interface{}
		// currNode  map[string]interface{}
	)

	// get process info
	processInfo, err := l.svcCtx.ProcessModel.FindOne(l.ctx, in.ProcessId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	// get process structure info
	err = json.Unmarshal([]byte(processInfo.Structure), &l.ProcessStructure.Structure)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	// get current node info
	// for _, node := range l.ProcessStructure.Structure["nodes"] {
	// 	if node["clazz"] == "start" {
	// 		currNode = node
	// 		break
	// 	}
	// }

	err = json.Unmarshal([]byte(in.State), &nodeState)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	nextNode, err := l.ProcessStructure.GetNode(nodeState[0].(map[string]interface{})["id"].(string))
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	for _, v := range template["form_data"] {
		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		l.FormData = append(l.FormData, data)
	}

	var condExprStatus bool

	switch nextNode["clazz"] {
	case "exclusiveGateway":
		edges, err := l.ProcessStructure.GetEdge(nextNode["id"].(string), "source")
		if err != nil {
			return nil, err
		}
	breakTag:
		for _, edge := range edges {
			edgeCondExpr := make([]map[string]interface{}, 0)
			err = json.Unmarshal([]byte(edge["conditionExpression"].(string)), &edgeCondExpr)
			if err != nil {
				return nil, err
			}
			for _, condExpr := range edgeCondExpr {
				// 条件判断
				condExprStatus, err = l.ConditionalJudgment(condExpr)
				if err != nil {
					return nil, err
				}
				if condExprStatus {
					// 进行节点跳转
					nodeValue, err := l.ProcessStructure.GetNode(edge["target"].(string))
					if err != nil {
						return nil, err
					}

					if nodeValue["clazz"] == "userTask" || nodeValue["clazz"] == "receiveTask" {
						if nodeValue["assignValue"] == nil || nodeValue["assignType"] == "" {
							err = errors.New("处理人不能为空")
							return nil, err
						}
					}
					nodeState[0].(map[string]interface{})["id"] = nodeValue["id"].(string)
					nodeState[0].(map[string]interface{})["label"] = nodeValue["label"]
					nodeState[0].(map[string]interface{})["processor"] = nodeValue["assignValue"]
					nodeState[0].(map[string]interface{})["process_method"] = nodeValue["assignType"]
					break breakTag
				}
			}
		}
		if !condExprStatus {
			err = errors.New("所有流转均不符合条件，请确认。")
			return nil, err
		}
	}

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

func (l *TicketAddLogic) ConditionalJudgment(condExpr map[string]interface{}) (result bool, err error) {
	var (
		condExprOk    bool
		condExprValue interface{}
	)

	defer func() {
		if r := recover(); r != nil {
			switch e := r.(type) {
			case string:
				err = errors.New(e)
			case error:
				err = e
			default:
				err = errors.New("未知错误")
			}
			return
		}
	}()

	for _, data := range l.FormData {
		var formData map[string]interface{}
		err = json.Unmarshal(data, &formData)
		if err != nil {
			return
		}
		if condExprValue, condExprOk = formData[condExpr["key"].(string)]; condExprOk {
			break
		}
	}

	if condExprValue == nil {
		err = errors.New("未查询到对应的表单数据。")
		return
	}

	// todo 待优化
	switch reflect.TypeOf(condExprValue).String() {
	case "string":
		switch condExpr["sign"] {
		case "==":
			if condExprValue.(string) == condExpr["value"].(string) {
				result = true
			}
		case "!=":
			if condExprValue.(string) != condExpr["value"].(string) {
				result = true
			}
		case ">":
			if condExprValue.(string) > condExpr["value"].(string) {
				result = true
			}
		case ">=":
			if condExprValue.(string) >= condExpr["value"].(string) {
				result = true
			}
		case "<":
			if condExprValue.(string) < condExpr["value"].(string) {
				result = true
			}
		case "<=":
			if condExprValue.(string) <= condExpr["value"].(string) {
				result = true
			}
		default:
			err = errors.New("目前仅支持6种常规判断类型，包括（等于、不等于、大于、大于等于、小于、小于等于）")
		}
	case "float64":
		switch condExpr["sign"] {
		case "==":
			if condExprValue.(float64) == condExpr["value"].(float64) {
				result = true
			}
		case "!=":
			if condExprValue.(float64) != condExpr["value"].(float64) {
				result = true
			}
		case ">":
			if condExprValue.(float64) > condExpr["value"].(float64) {
				result = true
			}
		case ">=":
			if condExprValue.(float64) >= condExpr["value"].(float64) {
				result = true
			}
		case "<":
			if condExprValue.(float64) < condExpr["value"].(float64) {
				result = true
			}
		case "<=":
			if condExprValue.(float64) <= condExpr["value"].(float64) {
				result = true
			}
		default:
			err = errors.New("目前仅支持6种常规判断类型，包括（等于、不等于、大于、大于等于、小于、小于等于）")
		}
	default:
		err = errors.New("条件判断目前仅支持字符串、整型。")
	}

	return
}
