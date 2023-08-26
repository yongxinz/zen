package process

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

type ProcessClassifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProcessClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessClassifyLogic {
	return &ProcessClassifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProcessClassifyLogic) ProcessClassify(req *types.ProcessClassifyReq) (resp *types.ProcessClassifyResp, err error) {
	res, err := l.svcCtx.WkfRpc.ProcessClassify(l.ctx, &wkfclient.ProcessClassifyReq{
		Name: req.Name,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get processclassify error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var process types.ProcessClassifyData
	var list []types.ProcessClassifyData
	for _, v := range res.Data {
		copier.Copy(&process, &v)
		list = append(list, process)
	}

	resp = &types.ProcessClassifyResp{
		List: list,
	}

	return
}
