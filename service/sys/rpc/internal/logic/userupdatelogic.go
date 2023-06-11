package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {
	sysUser, err := l.svcCtx.UserModel.FindOne(l.ctx, in.UserId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	_, err = l.svcCtx.DeptModel.FindOne(l.ctx, in.DeptId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.DeptIdErrorCode)
	}

	_, err = l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
	}

	err = copier.Copy(sysUser, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.UserModel.Update(l.ctx, sysUser)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.UserUpdateResp{}, nil
}
