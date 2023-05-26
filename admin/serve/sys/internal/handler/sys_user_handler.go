package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"serve/sys/internal/logic"
	"serve/sys/internal/svc"
	"serve/sys/internal/types"
)

// 处理登录
func loginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		sysUser := logic.NewSysUserLogic(r.Context(), ctx)
		resp, err := sysUser.Login(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
