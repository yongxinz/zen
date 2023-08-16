package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTemplateDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateDeleteLogic {
	return &TemplateDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TemplateDeleteLogic) TemplateDelete(in *wkf.TemplateDeleteReq) (*wkf.TemplateDeleteResp, error) {
	err := l.svcCtx.TemplateModel.Delete(l.ctx, in.TemplateId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TemplateDeleteResp{}, nil
}
