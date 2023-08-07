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

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *sys.MenuListReq) (*sys.MenuListResp, error) {
	menus, err := l.svcCtx.MenuModel.FindAll(l.ctx, 1, 1000)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("get menulist error: %s", err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var data []*sys.MenuListData
	for _, item := range menus {
		if item.ParentId.Int64 != 0 {
			continue
		}
		m := sys.MenuListData{
			MenuId:     item.Id,
			Title:      item.Title.String,
			MenuName:   item.MenuName.String,
			MenuType:   item.MenuType.String,
			Sort:       item.Sort.Int64,
			ParentId:   item.ParentId.Int64,
			Icon:       item.Icon.String,
			Path:       item.Path.String,
			Paths:      item.Paths,
			Permission: item.Permission.String,
			IsFrame:    item.IsFrame,
			NoCache:    item.NoCache.Int64,
			Action:     item.Action.String,
			Breadcrumb: item.Breadcrumb.String,
			Component:  item.Component.String,
			Visible:    item.Visible.String,
			CreateBy:   item.CreateBy.Int64,
			CreatedAt:  item.CreatedAt.Format(globalkey.SysDateFormat),
			UpdateBy:   item.UpdateBy.Int64,
			UpdatedAt:  item.UpdatedAt.Format(globalkey.SysDateFormat),
			Children:   []*sys.MenuListData{},
		}
		menu := menuListCall(menus, &m)
		data = append(data, menu)
	}

	return &sys.MenuListResp{
		Data: data,
	}, nil
}

func menuListCall(menuList []*model.SysMenu, menu *sys.MenuListData) *sys.MenuListData {
	list := menuList

	min := make([]*sys.MenuListData, 0)
	for j := 0; j < len(list); j++ {
		if menu.MenuId != list[j].ParentId.Int64 {
			continue
		}

		mi := sys.MenuListData{
			MenuId:     list[j].Id,
			Title:      list[j].Title.String,
			MenuName:   list[j].MenuName.String,
			MenuType:   list[j].MenuType.String,
			Sort:       list[j].Sort.Int64,
			ParentId:   list[j].ParentId.Int64,
			Icon:       list[j].Icon.String,
			Path:       list[j].Path.String,
			Paths:      list[j].Paths,
			Permission: list[j].Permission.String,
			IsFrame:    list[j].IsFrame,
			NoCache:    list[j].NoCache.Int64,
			Action:     list[j].Action.String,
			Breadcrumb: list[j].Breadcrumb.String,
			Component:  list[j].Component.String,
			Visible:    list[j].Visible.String,
			CreateBy:   list[j].CreateBy.Int64,
			CreatedAt:  list[j].CreatedAt.Format(globalkey.SysDateFormat),
			UpdateBy:   list[j].UpdateBy.Int64,
			UpdatedAt:  list[j].UpdatedAt.Format(globalkey.SysDateFormat),
			Children:   []*sys.MenuListData{},
		}
		if mi.MenuType != "F" {
			ms := menuListCall(menuList, &mi)
			min = append(min, ms)
		} else {
			min = append(min, &mi)
		}
	}
	menu.Children = min

	return menu
}
