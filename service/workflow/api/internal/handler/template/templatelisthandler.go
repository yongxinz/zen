package template

import (
	"net/http"

	"github.com/yongxin/zen/common/response"
	"github.com/yongxin/zen/service/workflow/api/internal/logic/template"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TemplateListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TemplateListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := template.NewTemplateListLogic(r.Context(), svcCtx)
		resp, err := l.TemplateList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, resp, err)
		}
	}
}
