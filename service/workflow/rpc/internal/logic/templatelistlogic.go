package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/common/globalkey"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type TemplateListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTemplateListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TemplateListLogic {
	return &TemplateListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TemplateListLogic) TemplateList(in *wkf.TemplateListReq) (*wkf.TemplateListResp, error) {
	res, err := l.svcCtx.TemplateModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get templatelist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.TemplateModel.Count(l.ctx)

	var data []*wkf.TemplateListData
	for _, item := range res {
		data = append(data, &wkf.TemplateListData{
			TemplateId:    item.Id,
			Name:          item.Name,
			FormStructure: item.FormStructure,
			Remark:        item.Remark,
			CreateBy:      item.CreateBy,
			UpdateBy:      item.UpdateBy,
			CreatedAt:     item.CreateAt.Format(globalkey.SysDateFormat),
			UpdatedAt:     item.UpdateAt.Format(globalkey.SysDateFormat),
		})
	}

	return &wkf.TemplateListResp{
		Count: count,
		Data:  data,
	}, nil
}
