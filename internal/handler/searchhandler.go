package handler

import (
	"net/http"

	"search/internal/logic"
	"search/internal/svc"
	"search/internal/types"
	"search/response"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SearchHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.ParsePath(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearchLogic(r.Context(), svcCtx)
		resp, err := l.Search(&req)
		//将返回结果交给统一返回请求处理
		response.Response(w, resp, err)

	}
}
