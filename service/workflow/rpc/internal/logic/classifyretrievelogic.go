package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClassifyRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyRetrieveLogic {
	return &ClassifyRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClassifyRetrieveLogic) ClassifyRetrieve(in *wkf.ClassifyRetrieveReq) (*wkf.ClassifyRetrieveResp, error) {
	res, err := l.svcCtx.ClassifyModel.FindOne(l.ctx, in.ClassifyId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ClassifyRetrieveResp{
		ClassifyId: res.Id,
		Name:       res.Name,
	}, nil
}
