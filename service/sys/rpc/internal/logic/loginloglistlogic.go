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

type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogListLogic) LoginLogList(in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	res, err := l.svcCtx.LoginLogModel.FindAll(l.ctx, in.Page, in.Size)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("get loginloglist failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	count, _ := l.svcCtx.LoginLogModel.Count(l.ctx)

	var data []*sys.LoginLogListData
	for _, item := range res {
		data = append(data, &sys.LoginLogListData{
			Username:      item.Username.String,
			Ipaddr:        item.Ipaddr.String,
			LoginLocation: item.LoginLocation.String,
			Browser:       item.Browser.String,
			Os:            item.Os.String,
			Msg:           item.Msg.String,
			LoginTime:     item.LoginTime.Format(globalkey.SysDateFormat),
		})
	}

	return &sys.LoginLogListResp{
		Count: count,
		Data:  data,
	}, nil
}
