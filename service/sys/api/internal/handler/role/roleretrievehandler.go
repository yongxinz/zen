package role

import (
	"net/http"

	"github.com/yongxin/zen/common/response"
	"github.com/yongxin/zen/service/sys/api/internal/logic/role"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/yongxin/zen/service/sys/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RoleRetrieveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleRetrieveReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleRetrieveLogic(r.Context(), svcCtx)
		resp, err := l.RoleRetrieve(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
