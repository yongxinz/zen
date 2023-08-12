package logic

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClassifyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyUpdateLogic {
	return &ClassifyUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClassifyUpdateLogic) ClassifyUpdate(in *wkf.ClassifyUpdateReq) (*wkf.ClassifyUpdateResp, error) {
	classify, err := l.svcCtx.ClassifyModel.FindOne(l.ctx, in.ClassifyId)
	if err != nil {
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	err = copier.Copy(classify, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	err = l.svcCtx.ClassifyModel.Update(l.ctx, classify)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ClassifyUpdateResp{}, nil
}
