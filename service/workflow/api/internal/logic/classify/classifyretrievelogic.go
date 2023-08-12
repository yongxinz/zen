package classify

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClassifyRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassifyRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyRetrieveLogic {
	return &ClassifyRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassifyRetrieveLogic) ClassifyRetrieve(req *types.ClassifyRetrieveReq) (resp *types.ClassifyRetrieveResp, err error) {
	res, err := l.svcCtx.WkfRpc.ClassifyRetrieve(l.ctx, &wkfclient.ClassifyRetrieveReq{
		ClassifyId: req.ClassifyId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ClassifyRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
