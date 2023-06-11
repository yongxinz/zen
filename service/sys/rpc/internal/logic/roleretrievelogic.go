package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleRetrieveLogic {
	return &RoleRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleRetrieveLogic) RoleRetrieve(in *sys.RoleRetrieveReq) (*sys.RoleRetrieveResp, error) {
	res, err := l.svcCtx.RoleModel.FindOne(l.ctx, in.RoleId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.RoleIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	rolemenu, err := l.svcCtx.RoleMenuModel.FindMenuIds(l.ctx, in.RoleId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var menuIds []int64
	for _, item := range rolemenu {
		menuIds = append(menuIds, item.MenuId)
	}

	return &sys.RoleRetrieveResp{
		RoleId:   res.Id,
		RoleKey:  res.RoleKey,
		RoleName: res.RoleName,
		Sort:     res.Sort,
		Status:   res.Status,
		Remark:   res.Remark,
		MenuIds:  menuIds,
	}, nil
}
