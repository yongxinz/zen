package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketProcessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketProcessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketProcessLogic {
	return &TicketProcessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type Ticket struct {
	*model.WkfTicket
	CurrentState string `json:"current_state"`
}

func (l *TicketProcessLogic) TicketProcess(in *wkf.TicketProcessReq) (*wkf.TicketProcessResp, error) {
	var (
		processStructure map[string]interface{}
		processNodes     []map[string]interface{}
		template         []*model.WkfTemplate
		formInfo         []*model.WkfForm
		ticketInfo       Ticket
		// formInfo         []*model.WkfForm
		stateList []map[string]interface{}
	)

	userId := int64(1)
	// userId, err := l.ctx.Value("userId").(json.Number).Int64()
	// fmt.Println("heeelo")
	// fmt.Println(err)
	// fmt.Println(userId)
	// fmt.Println("heeelo")
	// if err != nil {
	// 	return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	// }

	processInfo, err := l.svcCtx.ProcessModel.FindOne(l.ctx, in.ProcessId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	if processInfo.Structure != "" {
		err = json.Unmarshal([]byte(processInfo.Structure), &processStructure)
		if err != nil {
			err = fmt.Errorf("json转map失败，%v", err.Error())
			return nil, err
		}

		nodes := processStructure["nodes"].([]interface{})
		for i := 0; i < len(nodes); i++ {
			for j := 1; j < len(nodes)-i; j++ {
				s1 := nodes[j].(map[string]interface{})["sort"]
				s2 := nodes[j-1].(map[string]interface{})["sort"]
				if s1 == nil || s2 == nil {
					return nil, errors.New("流程未定义顺序属性，请确认")
				}
				left, _ := strconv.Atoi(s1.(string))
				right, _ := strconv.Atoi(s2.(string))
				if left < right {
					nodes[j], nodes[j-1] = nodes[j-1], nodes[j]
				}
			}
		}
		for _, node := range nodes {
			processNodes = append(processNodes, node.(map[string]interface{}))
		}
	}

	processInfoByte, _ := json.Marshal(processInfo)
	processNodesByte, _ := json.Marshal(processNodes)
	processEdgesByte, _ := json.Marshal(processStructure["edges"])

	// 流转信息
	circution, _ := l.svcCtx.CirculationModel.FindByTicket(l.ctx, in.TicketId)
	circutionByte, _ := json.Marshal(circution)

	if in.TicketId == 0 {
		var templateIds []int64
		err = json.Unmarshal([]byte(processInfo.Template), &templateIds)
		if err != nil {
			err = fmt.Errorf("json转map失败，%v", err.Error())
			return nil, err
		}
		template, _ = l.getTemplate(templateIds)
	} else {
		ticketInfo.WkfTicket, err = l.svcCtx.TicketModel.FindOne(l.ctx, in.TicketId)
		if err != nil {
			err = fmt.Errorf("查询工单信息失败，%v", err.Error())
			return nil, err
		}

		// 获取当前节点
		err = json.Unmarshal([]byte(ticketInfo.State), &stateList)
		if err != nil {
			err = fmt.Errorf("序列化节点列表失败，%v", err.Error())
			return nil, err
		}
		if len(stateList) == 0 {
			err = errors.New("当前工单没有下一节点数据")
			return nil, err
		}

		// 整理需要并行处理的数据
		if len(stateList) > 1 {
		continueHistoryTag:
			for _, v := range circution {
				status := false
				for i, s := range stateList {
					if v.Source == s["id"].(string) && v.Target != "" {
						status = true
						stateList = append(stateList[:i], stateList[i+1:]...)
						continue continueHistoryTag
					}
				}
				if !status {
					break
				}
			}
		}

		if len(stateList) > 0 {
		breakStateTag:
			for _, stateValue := range stateList {
				if processStructure["nodes"] != nil {
					for _, processNodeValue := range processStructure["nodes"].([]interface{}) {
						if stateValue["id"].(string) == processNodeValue.(map[string]interface{})["id"] {
							if _, ok := stateValue["processor"]; ok {
								for _, uId := range stateValue["processor"].([]interface{}) {
									if int64(uId.(float64)) == userId {
										ticketInfo.CurrentState = stateValue["id"].(string)
										break breakStateTag
									}
								}
							} else {
								err = errors.New("未查询到对应的处理人字段，请确认。")
								return nil, err
							}
						}
					}
				}
			}

			if ticketInfo.CurrentState == "" {
				ticketInfo.CurrentState = stateList[0]["id"].(string)
			}
		}

		formInfo, _ = l.svcCtx.FormModel.FindByTicket(l.ctx, in.TicketId)
	}
	templateByte, _ := json.Marshal(template)
	ticketByte, _ := json.Marshal(ticketInfo)
	formInfoByte, _ := json.Marshal(formInfo)

	return &wkf.TicketProcessResp{
		Process:     string(processInfoByte),
		Nodes:       string(processNodesByte),
		Edges:       string(processEdgesByte),
		Circulation: string(circutionByte),
		Template:    string(templateByte),
		FormData:    string(formInfoByte),
		Ticket:      string(ticketByte),
	}, nil
}

func (l *TicketProcessLogic) getTemplate(templateIds []int64) ([]*model.WkfTemplate, error) {
	template, err := l.svcCtx.TemplateModel.FindByIds(l.ctx, templateIds)
	if err != nil {
		err = fmt.Errorf("查询模板信息失败，%v", err.Error())
		return nil, err
	}
	return template, nil
}
