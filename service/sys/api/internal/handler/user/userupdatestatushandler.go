package user

import (
	"net/http"

	"github.com/yongxin/zen/common/response"
	"github.com/yongxin/zen/service/sys/api/internal/logic/user"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserUpdateStatusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserUpdateStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserUpdateStatusLogic(r.Context(), svcCtx)
		err := l.UserUpdateStatus(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			response.Response(w, nil, err)
		}
	}
}
