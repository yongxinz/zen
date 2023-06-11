package role

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleMenuTreeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleMenuTreeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleMenuTreeLogic {
	return &RoleMenuTreeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleMenuTreeLogic) RoleMenuTree(req *types.RoleMenuTreeReq) (resp *types.RoleMenuTreeResp, err error) {
	res, err := l.svcCtx.SysRpc.RoleMenuTree(l.ctx, &sysclient.RoleMenuTreeReq{
		RoleId: req.RoleId,
	})
	if err != nil {
		return nil, err
	}

	var menu types.RoleMenuTreeData
	var menus []types.RoleMenuTreeData
	for _, item := range res.Menus {
		copier.Copy(&menu, &item)
		menus = append(menus, menu)
	}
	resp = &types.RoleMenuTreeResp{
		Menus:       menus,
		CheckedKeys: res.CheckedKeys,
	}

	return
}
