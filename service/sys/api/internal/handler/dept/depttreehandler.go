package dept

import (
	"net/http"

	"github.com/yongxin/zen/common/response"
	"github.com/yongxin/zen/service/sys/api/internal/logic/dept"
	"github.com/yongxin/zen/service/sys/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func DeptTreeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := dept.NewDeptTreeLogic(r.Context(), svcCtx)
		resp, err := l.DeptTree()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
