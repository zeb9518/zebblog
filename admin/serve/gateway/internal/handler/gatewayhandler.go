package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"serve/gateway/internal/logic"
	"serve/gateway/internal/svc"
	"serve/gateway/internal/types"
)

func GatewayHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGatewayLogic(r.Context(), svcCtx)
		resp, err := l.Gateway(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}