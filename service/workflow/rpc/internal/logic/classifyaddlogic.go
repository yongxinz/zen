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

type ClassifyAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClassifyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyAddLogic {
	return &ClassifyAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClassifyAddLogic) ClassifyAdd(in *wkf.ClassifyAddReq) (*wkf.ClassifyAddResp, error) {
	var wkfClassify = new(model.WkfClassify)
	err := copier.Copy(wkfClassify, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.ClassifyModel.Insert(l.ctx, wkfClassify)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ClassifyAddResp{}, nil
}
