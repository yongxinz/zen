package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTemplateUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateUpdateLogic {
	return &TemplateUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TemplateUpdateLogic) TemplateUpdate(in *wkf.TemplateUpdateReq) (*wkf.TemplateUpdateResp, error) {
	template, err := l.svcCtx.TemplateModel.FindOne(l.ctx, in.TemplateId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(template, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.TemplateModel.Update(l.ctx, template)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TemplateUpdateResp{}, nil
}
