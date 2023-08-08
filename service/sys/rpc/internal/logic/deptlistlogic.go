package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/common/globalkey"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListLogic {
	return &DeptListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListLogic) DeptList(in *sys.DeptListReq) (*sys.DeptListResp, error) {
	depts, err := l.svcCtx.DeptModel.FindAll(l.ctx, 1, 1000)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get DeptList error: %s", err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var data []*sys.DeptListData
	for _, item := range depts {
		if item.ParentId.Int64 != 0 {
			continue
		}
		m := sys.DeptListData{
			DeptId:    item.Id,
			DeptPath:  item.DeptPath,
			DeptName:  item.DeptName.String,
			Leader:    item.Leader.Int64,
			Status:    item.Status.Int64,
			Sort:      item.Sort.Int64,
			ParentId:  item.ParentId.Int64,
			CreateBy:  item.CreateBy.Int64,
			CreatedAt: item.CreatedAt.Time.Format(globalkey.SysDateFormat),
			UpdateBy:  item.UpdateBy.Int64,
			UpdatedAt: item.UpdatedAt.Time.Format(globalkey.SysDateFormat),
			Children:  []*sys.DeptListData{},
		}
		deptInfo := deptListCall(depts, &m)
		data = append(data, deptInfo)
	}

	return &sys.DeptListResp{
		Data: data,
	}, nil
}

func deptListCall(deptList []*model.SysDept, dept *sys.DeptListData) *sys.DeptListData {
	list := deptList

	min := make([]*sys.DeptListData, 0)
	for j := 0; j < len(list); j++ {
		if dept.DeptId != list[j].ParentId.Int64 {
			continue
		}

		mi := sys.DeptListData{
			DeptId:    list[j].Id,
			DeptPath:  list[j].DeptPath,
			DeptName:  list[j].DeptName.String,
			Leader:    list[j].Leader.Int64,
			Status:    list[j].Status.Int64,
			Sort:      list[j].Sort.Int64,
			ParentId:  list[j].ParentId.Int64,
			CreateBy:  list[j].CreateBy.Int64,
			CreatedAt: list[j].CreatedAt.Time.Format(globalkey.SysDateFormat),
			UpdateBy:  list[j].UpdateBy.Int64,
			UpdatedAt: list[j].UpdatedAt.Time.Format(globalkey.SysDateFormat),
			Children:  []*sys.DeptListData{},
		}
		ms := deptListCall(deptList, &mi)
		min = append(min, ms)
	}
	dept.Children = min

	return dept
}
