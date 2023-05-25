package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"serve/internal/logic"
	"serve/internal/svc"
	"serve/internal/types"
)

func ServeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewServeLogic(r.Context(), svcCtx)
		resp, err := l.Serve(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
