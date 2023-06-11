package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostUpdateLogic {
	return &PostUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostUpdateLogic) PostUpdate(in *sys.PostUpdateReq) (*sys.PostUpdateResp, error) {
	sysPost, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.UserIdErrorCode)
	}

	err = copier.Copy(sysPost, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.PostModel.Update(l.ctx, sysPost)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.PostUpdateResp{}, nil
}
