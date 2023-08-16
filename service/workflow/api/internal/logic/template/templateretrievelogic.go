package template

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateRetrieveLogic {
	return &TemplateRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateRetrieveLogic) TemplateRetrieve(req *types.TemplateRetrieveReq) (resp *types.TemplateRetrieveResp, err error) {
	res, err := l.svcCtx.WkfRpc.TemplateRetrieve(l.ctx, &wkfclient.TemplateRetrieveReq{
		TemplateId: req.TemplateId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.TemplateRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
