package logic

import (
	"context"
	"encoding/json"
	"time"

	"github.com/yongxin/zen/common/captcha"
	"github.com/yongxin/zen/common/cryptx"
	"github.com/yongxin/zen/common/errorx"
	"github.com/yongxin/zen/common/jwtx"
	"github.com/yongxin/zen/service/sys/model"
	"github.com/yongxin/zen/service/sys/rpc/internal/svc"
	"github.com/yongxin/zen/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *sys.LoginRequest) (*sys.LoginResponse, error) {
	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil {
		if err == model.ErrNotFound {
			return nil, errorx.NewDefaultError(errorx.AccountErrorCode)
		}
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)
	if password != res.Password {
		return nil, errorx.NewDefaultError(errorx.PasswordErrorCode)
	}

	if !captcha.Verify(in.Uuid, in.Code, true) {
		return nil, errorx.NewDefaultError(errorx.CaptchaErrorCode)
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := jwtx.GetToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, res.Id)
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("gen jwt-token failed, params: %s, error: %s", reqStr, err.Error())
		return nil, errorx.NewSystemError(errorx.ServerErrorCode, err.Error())
	}

	return &sys.LoginResponse{
		CurrentAuthority: res.Username,
		Expire:           now + accessExpire,
		Token:            jwtToken,
	}, nil
}
