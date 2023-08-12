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

type ClassifyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClassifyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyListLogic {
	return &ClassifyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClassifyListLogic) ClassifyList(req *types.ClassifyListReq) (resp *types.ClassifyListResp, err error) {
	res, err := l.svcCtx.WkfRpc.ClassifyList(l.ctx, &wkfclient.ClassifyListReq{
		Page: req.PageIndex,
		Size: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get classifylist error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var classify types.ClassifyListData
	var list []types.ClassifyListData
	for _, v := range res.Data {
		copier.Copy(&classify, &v)
		list = append(list, classify)
	}

	resp = &types.ClassifyListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
