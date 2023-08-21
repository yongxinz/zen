package process

import (
	"context"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/yongxin/zen/service/workflow/rpc/wkfclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessRetrieveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessRetrieveLogic {
	return &ProcessRetrieveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessRetrieveLogic) ProcessRetrieve(req *types.ProcessRetrieveReq) (resp *types.ProcessRetrieveResp, err error) {
	res, err := l.svcCtx.WkfRpc.ProcessRetrieve(l.ctx, &wkfclient.ProcessRetrieveReq{
		ProcessId: req.ProcessId,
	})
	if err != nil {
		return nil, err
	}

	resp = &types.ProcessRetrieveResp{}
	err = copier.Copy(resp, res)

	return
}
