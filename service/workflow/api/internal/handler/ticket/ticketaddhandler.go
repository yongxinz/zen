package ticket

import (
	"net/http"

	"github.com/yongxin/zen/common/response"
	"github.com/yongxin/zen/service/workflow/api/internal/logic/ticket"
	"github.com/yongxin/zen/service/workflow/api/internal/svc"
	"github.com/yongxin/zen/service/workflow/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TicketAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TicketAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := ticket.NewTicketAddLogic(r.Context(), svcCtx)
		err := l.TicketAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			response.Response(w, nil, err)
		}
	}
}
