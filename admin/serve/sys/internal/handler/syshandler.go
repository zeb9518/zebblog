package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"sys/internal/logic"
	"sys/internal/svc"
	"sys/internal/types"
)

func SysHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewSysLogic(r.Context(), svcCtx)
		resp, err := l.Sys(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
