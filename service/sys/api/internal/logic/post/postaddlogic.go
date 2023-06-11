package post

import (
	"context"
	"encoding/json"

	"github.com/jinzhu/copier"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/yongxin/zen/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type PostAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPostAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PostAddLogic {
	return &PostAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PostAddLogic) PostAdd(req *types.PostAddReq) error {
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}
	req.CreateBy = userId
	req.UpdateBy = userId

	var sysPost sysclient.PostAddReq
	err = copier.Copy(&sysPost, req)
	if err != nil {
		return errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	_, err = l.svcCtx.SysRpc.PostAdd(l.ctx, &sysPost)
	if err != nil {
		return err
	}

	return nil
}
