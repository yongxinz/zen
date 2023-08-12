package classify

import (
	"net/http"

	"github.com/yongxin/zen/common/response"
	"github.com/yongxin/zen/service/workflow/api/internal/logic/classify"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ClassifyListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ClassifyListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := classify.NewClassifyListLogic(r.Context(), svcCtx)
		resp, err := l.ClassifyList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
