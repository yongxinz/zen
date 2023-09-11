package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/yongxin/zen/common/globalkey"
	"github.com/yongxin/zen/service/sys/rpc/sys"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TicketListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTicketListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TicketListLogic {
	return &TicketListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TicketListLogic) TicketList(in *wkf.TicketListReq) (*wkf.TicketListResp, error) {
	var (
		res  []*model.TicketList
		data []*wkf.TicketListData
		err  error
	)

	switch in.Category {
	case 1:
		userinfo, _ := l.svcCtx.SysRpc.UserRetrieve(l.ctx, &sys.UserInfoReq{
			UserId: in.UserId,
		})

		params := map[string]int64{
			"category": in.Category,
			"userId":   in.UserId,
			"roleId":   userinfo.RoleId,
			"deptId":   userinfo.DeptId,
		}
		res, _ = l.svcCtx.TicketModel.FindAll(l.ctx, in.Page, in.Size, params)
	case 2:
		params := map[string]int64{
			"category": in.Category,
			"userId":   in.UserId,
		}
		res, _ = l.svcCtx.TicketModel.FindAll(l.ctx, in.Page, in.Size, params)
	case 3:
		params := map[string]int64{
			"category": in.Category,
			"userId":   in.UserId,
		}
		res, _ = l.svcCtx.TicketModel.FindAll(l.ctx, in.Page, in.Size, params)
	case 4:
		params := map[string]int64{
			"category": in.Category,
			"userId":   in.UserId,
		}
		res, _ = l.svcCtx.TicketModel.FindAll(l.ctx, in.Page, in.Size, params)
	default:
		err = fmt.Errorf("参数错误，%v", err.Error())
		return nil, err
	}

	var stateList []map[string]interface{}
	for _, v := range res {
		err = json.Unmarshal([]byte(v.State), &stateList)
		if err != nil {
			err = fmt.Errorf("json反序列化失败，%v", err.Error())
			return nil, err
		}

		processorList := make([]int64, 0)
		for _, p := range stateList[0]["processor"].([]interface{}) {
			processorList = append(processorList, int64(p.(float64)))
		}
		stateName := stateList[0]["label"].(string)
		processMethod := stateList[0]["process_method"].(string)
		principals, _ := l.getPrincipal(processorList, processMethod)

		data = append(data, &wkf.TicketListData{
			TicketId:      v.Id,
			StateName:     stateName,
			ProcessName:   v.ProcessName,
			ProcessMethod: processMethod,
			Principals:    principals,
			IsEnd:         v.IsEnd,
			CreateBy:      v.CreateBy,
			UpdateBy:      v.UpdateBy,
			CreatedAt:     v.CreateAt.Format(globalkey.SysDateFormat),
			UpdatedAt:     v.UpdateAt.Format(globalkey.SysDateFormat),
		})
	}

	return &wkf.TicketListResp{
		Data: data,
	}, nil
}

func (l *TicketListLogic) getPrincipal(processor []int64, processMethod string) (string, error) {
	/*
		person 人员
		persongroup 人员组
		department 部门
		variable 变量
	*/
	var principalList []string

	switch processMethod {
	case "person":
		res, err := l.svcCtx.SysRpc.UserListByIds(l.ctx, &sys.UserListByIdsReq{
			UserIds: processor,
		})
		if err != nil {
			return "", err
		}
		for _, item := range res.List {
			principalList = append(principalList, item.UserName)
		}
	case "role":
		res, err := l.svcCtx.SysRpc.RoleListByIds(l.ctx, &sys.RoleListByIdsReq{
			RoleIds: processor,
		})
		if err != nil {
			return "", err
		}
		for _, item := range res.Data {
			principalList = append(principalList, item.RoleName)
		}
	case "department":
		res, err := l.svcCtx.SysRpc.DeptListByIds(l.ctx, &sys.DeptListByIdsReq{
			DeptIds: processor,
		})
		if err != nil {
			return "", err
		}
		for _, item := range res.Data {
			principalList = append(principalList, item.DeptName)
		}
	case "variable":
		for _, p := range processor {
			switch p {
			case 1:
				principalList = append(principalList, "创建者")
			case 2:
				principalList = append(principalList, "创建者负责人")
			}
		}
	}
	return strings.Join(principalList, ","), nil
}
