package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuRoleLogic {
	return &MenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuRoleLogic) MenuRole(in *sys.MenuRoleReq) (*sys.MenuRoleResp, error) {
	var err error
	var menus []*model.SysMenu

	roleInfo, _ := l.svcCtx.RoleModel.FindOne(l.ctx, in.RoleId)
	if roleInfo.RoleKey == "admin" {
		count, _ := l.svcCtx.MenuModel.Count(l.ctx)
		menus, err = l.svcCtx.MenuModel.FindAll(l.ctx, 1, count)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
	} else {
		res, err := l.svcCtx.RoleMenuModel.FindMenuIds(l.ctx, in.RoleId)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}

		var menuIds []int64
		for _, item := range res {
			menuIds = append(menuIds, item.MenuId)
		}

		menus, err = l.svcCtx.MenuModel.FindMenuList(l.ctx, menuIds)
		if err != nil {
			return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
		}
	}

	data := make([]*sys.MenuTree, 0)
	for _, menu := range menus {
		if menu.ParentId.Int64 != 0 {
			continue
		}
		m := sys.MenuTree{
			Data: &sys.MenuListData{
				MenuId:     menu.Id,
				MenuName:   menu.MenuName.String,
				MenuType:   menu.MenuType.String,
				Title:      menu.Title.String,
				Icon:       menu.Icon.String,
				Path:       menu.Path.String,
				Paths:      menu.Paths,
				Action:     menu.Action.String,
				Permission: menu.Permission.String,
				ParentId:   menu.ParentId.Int64,
				Breadcrumb: menu.Breadcrumb.String,
				Component:  menu.Component.String,
				NoCache:    menu.NoCache.Int64,
				Sort:       menu.Sort.Int64,
				Visible:    menu.Visible.String,
			},
			Children: []*sys.MenuTree{},
		}
		menuInfo := menuCall(menus, &m)
		data = append(data, menuInfo)
	}

	return &sys.MenuRoleResp{
		Data: data,
	}, nil
}

func menuCall(menuList []*model.SysMenu, menu *sys.MenuTree) *sys.MenuTree {
	list := menuList

	min := make([]*sys.MenuTree, 0)
	for j := 0; j < len(list); j++ {
		if menu.Data.MenuId != list[j].ParentId.Int64 {
			continue
		}

		mi := sys.MenuTree{
			Data:     &sys.MenuListData{},
			Children: []*sys.MenuTree{},
		}
		mi.Data.MenuId = list[j].Id
		mi.Data.MenuName = list[j].MenuName.String
		mi.Data.MenuType = list[j].MenuType.String
		mi.Data.Title = list[j].Title.String
		mi.Data.Icon = list[j].Icon.String
		mi.Data.Path = list[j].Path.String
		mi.Data.Paths = list[j].Paths
		mi.Data.Action = list[j].Action.String
		mi.Data.Permission = list[j].Permission.String
		mi.Data.ParentId = list[j].ParentId.Int64
		mi.Data.NoCache = list[j].NoCache.Int64
		mi.Data.Breadcrumb = list[j].Breadcrumb.String
		mi.Data.Component = list[j].Component.String
		mi.Data.Sort = list[j].Sort.Int64
		mi.Data.Visible = list[j].Visible.String
		mi.Children = []*sys.MenuTree{}

		if mi.Data.MenuType != "F" {
			ms := menuCall(menuList, &mi)
			min = append(min, ms)
		} else {
			min = append(min, &mi)
		}
	}
	menu.Children = min

	return menu
}
