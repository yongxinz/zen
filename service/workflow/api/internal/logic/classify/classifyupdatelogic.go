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

type ClassifyUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassifyUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyUpdateLogic {
	return &ClassifyUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassifyUpdateLogic) ClassifyUpdate(req *types.ClassifyUpdateReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.UpdateBy = userId

	var classify wkfclient.ClassifyUpdateReq
	err = copier.Copy(&classify, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.WkfRpc.ClassifyUpdate(l.ctx, &classify)
	if err != nil {
		return err
	}

	return nil
}
