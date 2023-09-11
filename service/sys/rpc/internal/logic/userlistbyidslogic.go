package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListByIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListByIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListByIdsLogic {
	return &UserListByIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListByIdsLogic) UserListByIds(in *sys.UserListByIdsReq) (*sys.UserListResp, error) {
	res, err := l.svcCtx.UserModel.FindByIds(l.ctx, in.UserIds)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("FindByIds failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	var data []*sys.UserListData
	for _, user := range res {
		data = append(data, &sys.UserListData{
			UserId:   user.Id,
			UserName: user.Username,
		})
	}

	return &sys.UserListResp{
		Count: 0,
		List:  data,
	}, nil
}
