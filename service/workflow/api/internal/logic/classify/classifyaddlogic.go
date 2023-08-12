package classify

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassifyAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyAddLogic {
	return &ClassifyAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassifyAddLogic) ClassifyAdd(req *types.ClassifyAddReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.CreateBy = userId
	req.UpdateBy = userId

	var classify wkfclient.ClassifyAddReq
	err = copier.Copy(&classify, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.WkfRpc.ClassifyAdd(l.ctx, &classify)
	if err != nil {
		return err
	}

	return nil
}
