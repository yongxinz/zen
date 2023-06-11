package role

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

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleListReq) (resp *types.RoleListResp, err error) {
	res, err := l.svcCtx.SysRpc.RoleList(l.ctx, &sysclient.RoleListReq{
		Page: req.PageIndex,
		Size: req.PageSize,
	})

	if err != nil {
		data, _ := json.Marshal(req)
		logx.WithContext(l.ctx).Errorf("get rolelist error: %s, params: %s", err.Error(), string(data))
		return nil, errorx.NewDefaultError(errorx.ServerErrorCode)
	}

	var post types.RoleListData
	var list []types.RoleListData
	for _, v := range res.Data {
		copier.Copy(&post, &v)
		list = append(list, post)
	}

	resp = &types.RoleListResp{
		List: list,
		Pagination: types.Pagination{
			PageIndex: req.PageIndex,
			PageSize:  req.PageSize,
			Count:     res.Count,
		},
	}

	return
}
