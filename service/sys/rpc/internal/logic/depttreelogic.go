package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptTreeLogic {
	return &DeptTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptTreeLogic) DeptTree(in *sys.DeptTreeReq) (*sys.DeptTreeResp, error) {
	depts, err := l.svcCtx.DeptModel.FindAll(l.ctx, 1, 1000)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get DeptTree error: %s", err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var data []*sys.DeptTreeData
	for _, item := range depts {
		if item.ParentId.Int64 != 0 {
			continue
		}
		m := sys.DeptTreeData{
			Id:       item.Id,
			Label:    item.DeptName.String,
			Children: []*sys.DeptTreeData{},
		}
		deptInfo := deptTreeCall(depts, &m)
		data = append(data, deptInfo)
	}

	return &sys.DeptTreeResp{
		Data: data,
	}, nil
}

func deptTreeCall(deptList []*model.SysDept, dept *sys.DeptTreeData) *sys.DeptTreeData {
	list := deptList

	min := make([]*sys.DeptTreeData, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId.Int64 {
			continue
		}

		mi := sys.DeptTreeData{
			Id:       list[j].Id,
			Label:    list[j].DeptName.String,
			Children: []*sys.DeptTreeData{},
		}
		ms := deptTreeCall(deptList, &mi)
		min = append(min, ms)
	}
	dept.Children = min

	return dept
}
