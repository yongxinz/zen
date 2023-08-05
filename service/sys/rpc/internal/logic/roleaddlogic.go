package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *sys.RoleAddReq) (*sys.RoleAddResp, error) {
	var SysRole = new(model.SysRole)
	err := copier.Copy(SysRole, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	res, err := l.svcCtx.RoleModel.Insert(l.ctx, SysRole)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	roleId, _ := res.LastInsertId()
	l.svcCtx.RoleMenuModel.DeleteByRoleId(l.ctx, roleId)
	for _, menuId := range in.MenuIds {
		l.svcCtx.RoleMenuModel.Insert(l.ctx, &model.SysRoleMenu{RoleId: roleId, MenuId: menuId})
	}

	return &sys.RoleAddResp{}, nil
}
