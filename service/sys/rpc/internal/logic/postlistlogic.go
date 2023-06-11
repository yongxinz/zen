package logic

import (
	"context"
	"encoding/json"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/common/globalkey"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostListLogic {
	return &PostListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostListLogic) PostList(in *sys.PostListReq) (*sys.PostListResp, error) {
	res, err := l.svcCtx.PostModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get postlist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.PostModel.Count(l.ctx)

	var data []*sys.PostListData
	for _, item := range res {
		data = append(data, &sys.PostListData{
			PostId:    item.Id,
			PostName:  item.PostName,
			PostCode:  item.PostCode,
			Sort:      item.Sort,
			Status:    item.Status,
			CreateBy:  item.CreateBy,
			UpdateBy:  item.UpdateBy,
			CreatedAt: item.CreatedAt.Format(globalkey.SysDateFormat),
			UpdatedAt: item.UpdatedAt.Format(globalkey.SysDateFormat),
		})
	}

	return &sys.PostListResp{
		Count: count,
		Data:  data,
	}, nil
}
