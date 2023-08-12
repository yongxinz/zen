package classify

import (
	"context"

	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassifyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyDeleteLogic {
	return &ClassifyDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassifyDeleteLogic) ClassifyDelete(req *types.ClassifyDeleteReq) error {
	_, err := l.svcCtx.WkfRpc.ClassifyDelete(l.ctx, &wkfclient.ClassifyDeleteReq{
		ClassifyId: req.ClassifyId,
	})

	return err
}
