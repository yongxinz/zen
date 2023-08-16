package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTemplateRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateRetrieveLogic {
	return &TemplateRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TemplateRetrieveLogic) TemplateRetrieve(in *wkf.TemplateRetrieveReq) (*wkf.TemplateRetrieveResp, error) {
	res, err := l.svcCtx.TemplateModel.FindOne(l.ctx, in.TemplateId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.TemplateRetrieveResp{
		TemplateId:    res.Id,
		Name:          res.Name,
		FormStructure: res.FormStructure,
		Remark:        res.Remark,
	}, nil
}
