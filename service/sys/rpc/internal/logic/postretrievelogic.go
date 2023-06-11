package logic

import (
	"context"

	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostRetrieveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPostRetrieveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostRetrieveLogic {
	return &PostRetrieveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PostRetrieveLogic) PostRetrieve(in *sys.PostRetrieveReq) (*sys.PostRetrieveResp, error) {
	res, err := l.svcCtx.PostModel.FindOne(l.ctx, in.PostId)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.PostIdErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.PostRetrieveResp{
		PostId:   res.Id,
		PostName: res.PostName,
		PostCode: res.PostCode,
		Sort:     res.Sort,
		Status:   res.Status,
		Remark:   res.Remark,
	}, nil
}
