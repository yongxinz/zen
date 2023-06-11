package post

import (
	"context"

	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostDeleteLogic {
	return &PostDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostDeleteLogic) PostDelete(req *types.PostDeleteReq) error {
	_, err := l.svcCtx.SysRpc.PostDelete(l.ctx, &sysclient.PostDeleteReq{
		PostIds: req.Ids,
	})
	return err
}
