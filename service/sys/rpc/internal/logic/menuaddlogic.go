package logic

import (
	"context"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *sys.MenuAddReq) (*sys.MenuAddResp, error) {
	var SysMenu = new(model.SysMenu)
	err := copier.Copy(SysMenu, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	res, err := l.svcCtx.MenuModel.Insert(l.ctx, SysMenu)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	menuId, _ := res.LastInsertId()
	SysMenu, _ = l.svcCtx.MenuModel.FindOne(l.ctx, menuId)
	parentId := SysMenu.ParentId.Int64

	var deptPath string
	if parentId == 0 {
		deptPath = "/0/" + strconv.FormatInt(SysMenu.Id, 10) + "/"
	} else {
		var parentDept = new(model.SysMenu)
		parentDept, _ = l.svcCtx.MenuModel.FindOne(l.ctx, SysMenu.ParentId.Int64)
		deptPath = parentDept.Paths + strconv.FormatInt(SysMenu.Id, 10) + "/"
	}

	SysMenu.Paths = deptPath
	err = l.svcCtx.MenuModel.Update(l.ctx, SysMenu)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.MenuAddResp{}, nil
}
