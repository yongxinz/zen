package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeptListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeptListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeptListByIdsLogic {
	return &DeptListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeptListByIdsLogic) DeptListByIds(in *sys.DeptListByIdsReq) (*sys.DeptListResp, error) {
	depts, err := l.svcCtx.DeptModel.FindByIds(l.ctx, in.DeptIds)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("FindByIds error: %s", err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var data []*sys.DeptListData
	for _, item := range depts {
		m := &sys.DeptListData{
			DeptId:   item.Id,
			DeptName: item.DeptName.String,
		}
		data = append(data, m)
	}

	return &sys.DeptListResp{
		Data: data,
	}, nil
}
