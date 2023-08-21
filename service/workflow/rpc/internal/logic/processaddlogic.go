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

type ProcessAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessAddLogic {
	return &ProcessAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcessAddLogic) ProcessAdd(in *wkf.ProcessAddReq) (*wkf.ProcessAddResp, error) {
	var wkfProcess = new(model.WkfProcess)
	err := copier.Copy(wkfProcess, in)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.ProcessModel.Insert(l.ctx, wkfProcess)
	if err != nil {
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &wkf.ProcessAddResp{}, nil
}
