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

type ClassifyListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClassifyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClassifyListLogic {
	return &ClassifyListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClassifyListLogic) ClassifyList(in *wkf.ClassifyListReq) (*wkf.ClassifyListResp, error) {
	res, err := l.svcCtx.ClassifyModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get classifylist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.ClassifyModel.Count(l.ctx)

	var data []*wkf.ClassifyListData
	for _, item := range res {
		data = append(data, &wkf.ClassifyListData{
			ClassifyId: item.Id,
			Name:       item.Name,
			CreateBy:   item.CreateBy,
			UpdateBy:   item.UpdateBy,
			CreatedAt:  item.CreateAt.Format(globalkey.SysDateFormat),
			UpdatedAt:  item.UpdateAt.Format(globalkey.SysDateFormat),
		})
	}

	return &wkf.ClassifyListResp{
		Count: count,
		Data:  data,
	}, nil
}
