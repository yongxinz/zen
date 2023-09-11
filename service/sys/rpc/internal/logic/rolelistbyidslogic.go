package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListByIdsLogic {
	return &RoleListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListByIdsLogic) RoleListByIds(in *sys.RoleListByIdsReq) (*sys.RoleListResp, error) {
	res, err := l.svcCtx.RoleModel.FindByIds(l.ctx, in.RoleIds)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("FindByIds failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var data []*sys.RoleListData
	for _, item := range res {
		data = append(data, &sys.RoleListData{
			RoleId:   item.Id,
			RoleName: item.RoleName,
		})
	}

	return &sys.RoleListResp{
		Count: 0,
		Data:  data,
	}, nil
}
