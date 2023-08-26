package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/workflow/model"
	"github.com/yongxin/zen/service/workflow/rpc/internal/svc"
	"github.com/yongxin/zen/service/workflow/rpc/wkf"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProcessClassifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProcessClassifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProcessClassifyLogic {
	return &ProcessClassifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProcessClassifyLogic) ProcessClassify(in *wkf.ProcessClassifyReq) (*wkf.ProcessClassifyResp, error) {
	res, err := l.svcCtx.ProcessModel.FindProcessClassify(l.ctx, in.Name)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get processclassify failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	mp := make(map[int64][]*model.WkfProcessClassify)
	for _, item := range res {
		mp[item.ClassifyId] = append(mp[item.ClassifyId], item)
	}

	var data []*wkf.ProcessClassifyData
	for classifyId, process := range mp {
		processData := []*wkf.ProcessListData{}
		for _, item := range process {
			processData = append(processData, &wkf.ProcessListData{
				ProcessId: item.Id,
				Name:      item.Name,
				Icon:      item.Icon,
				Remark:    item.Remark,
			})
		}
		data = append(data, &wkf.ProcessClassifyData{
			ClassifyId: classifyId,
			Name:       process[0].ClassifyName,
			Process:    processData,
		})
	}

	return &wkf.ProcessClassifyResp{
		Data: data,
	}, nil
}
