package handler

import (
	"net/http"

	"blog/internal/logic"
	"blog/internal/svc"
	"blog/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewBlogLogic(r.Context(), svcCtx)
		resp, err := l.Blog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
