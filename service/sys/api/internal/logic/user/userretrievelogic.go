package user

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRetrieveLogic {
	return &UserRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRetrieveLogic) UserRetrieve(req *types.UserRetrieveReq) (resp *types.UserRetrieveResp, err error) {
	res, err := l.svcCtx.SysRpc.UserRetrieve(l.ctx, &sysclient.UserInfoReq{
		UserId: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.UserRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
