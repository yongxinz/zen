package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTemplateAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateAddLogic {
	return &TemplateAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TemplateAddLogic) TemplateAdd(in *wkf.TemplateAddReq) (*wkf.TemplateAddResp, error) {
	var wkfTemplate = new(model.WkfTemplate)
	err := copier.Copy(wkfTemplate, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.TemplateModel.Insert(l.ctx, wkfTemplate)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TemplateAddResp{}, nil
}
