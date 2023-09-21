package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/yongxin/zen/service/sys/rpc/sys"
	"github.com/yongxin/zen/service/workflow/model"
)

type UserPermission struct {
	TicketInfo *model.WkfTicket
	UserInfo   *sys.UserRetrieveResp
}

// Determine whether the user has operating permissions for the current node
func (l *UserPermission) JudgeUserPermission() (bool, error) {
	var currentState map[string]interface{}

	// Get the current node that the user needs to process
	err := json.Unmarshal([]byte(l.TicketInfo.State), &currentState)
	if err != nil {
		err = fmt.Errorf("JudgeUserPermission Ticket State Unmarshal, %v", err.Error())
		return false, err
	}

	switch currentState["process_method"].(string) {
	case "person":
		for _, processorValue := range currentState["processor"].([]interface{}) {
			if int64(processorValue.(float64)) == l.UserInfo.UserId {
				return true, nil
			}
		}
	case "role":
		for _, processorValue := range currentState["processor"].([]interface{}) {
			if int64(processorValue.(float64)) == l.UserInfo.RoleId {
				return true, nil
			}
		}
	case "department":
		for _, processorValue := range currentState["processor"].([]interface{}) {
			if int64(processorValue.(float64)) == l.UserInfo.DeptId {
				return true, nil
			}
		}
	}

	return false, errors.New("no permission to process this node")
}

type ProcessStructure struct {
	Structure map[string][]map[string]interface{}
}

// get node info
func (p *ProcessStructure) GetNode(stateId string) (nodeValue map[string]interface{}, err error) {
	for _, node := range p.Structure["nodes"] {
		if node["id"] == stateId {
			nodeValue = node
		}
	}
	return
}

// get edge info
func (p *ProcessStructure) GetEdge(stateId string, classify string) (edgeValue []map[string]interface{}, err error) {
	for _, edge := range p.Structure["edges"] {
		if edge[classify] == stateId {
			edgeValue = append(edgeValue, edge)
		}
	}

	var left, right int

	for i := 0; i < len(edgeValue)-1; i++ {
		for j := i + 1; j < len(edgeValue); j++ {
			if t, ok := edgeValue[i]["sort"]; ok {
				left, _ = strconv.Atoi(t.(string))
			}
			if t, ok := edgeValue[j]["sort"]; ok {
				right, _ = strconv.Atoi(t.(string))
			}
			if left > right {
				edgeValue[j], edgeValue[i] = edgeValue[i], edgeValue[j]
			}
		}
	}

	return
}
