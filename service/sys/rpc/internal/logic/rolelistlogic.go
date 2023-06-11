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

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	res, err := l.svcCtx.RoleModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get rolelist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.RoleModel.Count(l.ctx)

	var data []*sys.RoleListData
	for _, item := range res {
		data = append(data, &sys.RoleListData{
			RoleId:    item.Id,
			RoleName:  item.RoleName,
			RoleKey:   item.RoleKey,
			Sort:      item.Sort,
			Remark:    item.Remark,
			Status:    item.Status,
			CreateBy:  item.CreateBy,
			UpdateBy:  item.UpdateBy,
			CreatedAt: item.CreatedAt.Format(globalkey.SysDateFormat),
			UpdatedAt: item.UpdatedAt.Format(globalkey.SysDateFormat),
		})
	}

	return &sys.RoleListResp{
		Count: count,
		Data:  data,
	}, nil
}
