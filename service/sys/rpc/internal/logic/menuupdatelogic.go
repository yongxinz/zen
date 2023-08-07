package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	sysMenu, err := l.svcCtx.MenuModel.FindOne(l.ctx, in.MenuId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.PermMenuIdErrorCode)
	}

	err = copier.Copy(sysMenu, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.MenuModel.Update(l.ctx, sysMenu)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.MenuUpdateResp{}, nil
}
