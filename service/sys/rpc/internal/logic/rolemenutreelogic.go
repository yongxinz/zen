package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuTreeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuTreeLogic {
	return &RoleMenuTreeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleMenuTreeLogic) RoleMenuTree(in *sys.RoleMenuTreeReq) (*sys.RoleMenuTreeResp, error) {
	menus, _ := l.svcCtx.MenuModel.FindAll(l.ctx, 1, 1000)

	m := make([]*sys.RoleMenuTreeData, 0)
	for i := 0; i < len(menus); i++ {
		if menus[i].ParentId.Int64 != 0 {
			continue
		}
		e := sys.RoleMenuTreeData{}
		e.Id = menus[i].Id
		e.Label = menus[i].Title.String
		deptsInfo := menuLabelCall(menus, &e)

		m = append(m, deptsInfo)
	}

	menuIds := make([]int64, 0)
	if in.RoleId != 0 {
		res, err := l.svcCtx.RoleMenuModel.FindMenuIds(l.ctx, in.RoleId)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
		for _, item := range res {
			menuIds = append(menuIds, item.MenuId)
		}
	}

	return &sys.RoleMenuTreeResp{
		Menus:       m,
		CheckedKeys: menuIds,
	}, nil
}

func menuLabelCall(eList []*model.SysMenu, menu *sys.RoleMenuTreeData) *sys.RoleMenuTreeData {
	list := eList

	min := make([]*sys.RoleMenuTreeData, 0)
	for j := 0; j < len(list); j++ {

		if menu.Id != list[j].ParentId.Int64 {
			continue
		}
		mi := sys.RoleMenuTreeData{}
		mi.Id = list[j].Id
		mi.Label = list[j].Title.String
		mi.Children = []*sys.RoleMenuTreeData{}
		if list[j].MenuType.String != "F" {
			ms := menuLabelCall(eList, &mi)
			min = append(min, ms)
		} else {
			min = append(min, &mi)
		}
	}
	if len(min) > 0 {
		menu.Children = min
	} else {
		menu.Children = nil
	}
	return menu
}
