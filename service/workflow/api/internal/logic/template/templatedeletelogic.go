package template

import (
	"context"

	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateDeleteLogic {
	return &TemplateDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateDeleteLogic) TemplateDelete(req *types.TemplateDeleteReq) error {
	_, err := l.svcCtx.WkfRpc.TemplateDelete(l.ctx, &wkfclient.TemplateDeleteReq{
		TemplateId: req.TemplateId,
	})

	return err
}
