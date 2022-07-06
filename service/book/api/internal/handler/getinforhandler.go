package handler

import (
	"net/http"

	"go-zero-bookmall/service/book/api/internal/logic"
	"go-zero-bookmall/service/book/api/internal/svc"
	"go-zero-bookmall/service/book/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetInfoLogic(r.Context(), svcCtx)
		resp, err := l.GetInfo(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
