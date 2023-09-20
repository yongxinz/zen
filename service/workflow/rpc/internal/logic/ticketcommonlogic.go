package logic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/yongxin/zen/service/sys/rpc/sys"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
)

// Determine whether the user has operating permissions for the current node
func JudgeUserPermission(ctx context.Context, ticketId, userId int64) (bool, error) {
	svcCtx := &svc.ServiceContext{}
	var currentState map[string]interface{}

	fmt.Println(11111111)
	// search ticket info
	ticketInfo, err := svcCtx.TicketModel.FindOne(ctx, ticketId)
	if err != nil {
		err = fmt.Errorf("JudgeUserPermission Ticket FindOne, %v", err.Error())
		return false, err
	}
	fmt.Println(2222222)

	// Get the current node that the user needs to process
	err = json.Unmarshal([]byte(ticketInfo.State), &currentState)
	if err != nil {
		err = fmt.Errorf("JudgeUserPermission Ticket State Unmarshal, %v", err.Error())
		return false, err
	}

	// search process info
	// process, err := svcCtx.ProcessModel.FindOne(ctx, ticketInfo.ProcessId)
	// if err != nil {
	// 	err = fmt.Errorf("JudgeUserPermission Process FindOne, %v", err.Error())
	// 	return false, err
	// }

	// search user info
	userInfo, err := svcCtx.SysRpc.UserRetrieve(ctx, &sys.UserInfoReq{
		UserId: userId,
	})
	if err != nil {
		return false, err
	}

	switch currentState["process_method"].(string) {
	case "person":
		for _, processorValue := range currentState["processor"].([]interface{}) {
			if int64(processorValue.(float64)) == userId {
				return true, nil
			}
		}
	case "role":
		for _, processorValue := range currentState["processor"].([]interface{}) {
			if int64(processorValue.(float64)) == userInfo.RoleId {
				return true, nil
			}
		}
	case "department":
		for _, processorValue := range currentState["processor"].([]interface{}) {
			if int64(processorValue.(float64)) == userInfo.DeptId {
				return true, nil
			}
		}
	}

	return false, errors.New("no permission to process this node")
}
