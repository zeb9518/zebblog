// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"serve/sys/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/sys/login",
				Handler: loginHandler(serverCtx),
			},
		},
	)
}