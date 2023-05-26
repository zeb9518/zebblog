package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"serve/blog/internal/logic"
	"serve/blog/internal/svc"
	"serve/blog/internal/types"
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
