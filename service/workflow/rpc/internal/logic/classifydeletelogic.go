package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClassifyDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyDeleteLogic {
	return &ClassifyDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClassifyDeleteLogic) ClassifyDelete(in *wkf.ClassifyDeleteReq) (*wkf.ClassifyDeleteResp, error) {
	err := l.svcCtx.ClassifyModel.Delete(l.ctx, in.ClassifyId)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ClassifyDeleteResp{}, nil
}
