package template

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

type TemplateListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateListLogic {
	return &TemplateListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TemplateListLogic) TemplateList(req *types.TemplateListReq) (resp *types.TemplateListResp, err error) {
	res, err := l.svcCtx.WkfRpc.TemplateList(l.ctx, &wkfclient.TemplateListReq{
		Page: req.PageIndex,
		Size: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get templatelist error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var template types.TemplateListData
	var list []types.TemplateListData
	for _, v := range res.Data {
		copier.Copy(&template, &v)
		list = append(list, template)
	}

	resp = &types.TemplateListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
