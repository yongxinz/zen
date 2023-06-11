package user

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResp, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	res, err := l.svcCtx.SysRpc.UserInfo(l.ctx, &sysclient.UserInfoReq{
		UserId: userId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserInfoResp{
		Avatar:       res.Avatar,
		Introduction: res.Introduction,
		Name:         res.Name,
		UserName:     res.UserName,
		UserId:       res.UserId,
		DeptId:       res.DeptId,
		Buttons:      res.Buttons,
		Roles:        res.Roles,
		Permissions:  res.Permissions,
	}

	return
}
