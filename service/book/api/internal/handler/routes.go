// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go-zero-bookmall/service/book/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/book/add",
				Handler: AddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/book/getprofile",
				Handler: GetInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/book/update",
				Handler: UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/book/search",
				Handler: SearchHandler(serverCtx),
			},
		},
	)
}