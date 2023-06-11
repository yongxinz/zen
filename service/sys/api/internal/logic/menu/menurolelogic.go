package menu

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuRoleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuRoleLogic {
	return &MenuRoleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuRoleLogic) MenuRole() (resp []*types.MenuRoleResp, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	userinfo, err := l.svcCtx.SysRpc.UserInfo(l.ctx, &sysclient.UserInfoReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	res, _ := l.svcCtx.SysRpc.MenuRole(l.ctx, &sysclient.MenuRoleReq{
		RoleId: userinfo.RoleId,
	})

	for _, menu := range res.Data {
		children := make([]types.MenuRoleResp, 0)
		for _, child := range menu.Children {
			children = append(children, types.MenuRoleResp{
				MenuId:     child.Data.MenuId,
				MenuName:   child.Data.MenuName,
				MenuType:   child.Data.MenuType,
				Title:      child.Data.Title,
				Path:       child.Data.Path,
				Paths:      child.Data.Paths,
				Icon:       child.Data.Icon,
				Breadcrumb: child.Data.Breadcrumb,
				Component:  child.Data.Component,
				Action:     child.Data.Action,
				IsFrame:    child.Data.IsFrame,
				ParentId:   child.Data.ParentId,
				NoCache:    child.Data.NoCache,
				Sort:       child.Data.Sort,
				Visible:    child.Data.Visible,
			})
		}
		resp = append(resp, &types.MenuRoleResp{
			MenuId:     menu.Data.MenuId,
			MenuName:   menu.Data.MenuName,
			MenuType:   menu.Data.MenuType,
			Title:      menu.Data.Title,
			Path:       menu.Data.Path,
			Paths:      menu.Data.Paths,
			Icon:       menu.Data.Icon,
			Breadcrumb: menu.Data.Breadcrumb,
			Component:  menu.Data.Component,
			Action:     menu.Data.Action,
			IsFrame:    menu.Data.IsFrame,
			ParentId:   menu.Data.ParentId,
			NoCache:    menu.Data.NoCache,
			Sort:       menu.Data.Sort,
			Visible:    menu.Data.Visible,
			Children:   children,
		})
	}

	return
}
