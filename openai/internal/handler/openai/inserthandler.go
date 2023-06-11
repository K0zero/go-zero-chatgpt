package openai

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-chatgpt/openai/internal/logic/openai"
	"go-zero-chatgpt/openai/internal/svc"
	"go-zero-chatgpt/openai/internal/types"
)

func InsertHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertAPIReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := openai.NewInsertLogic(r.Context(), svcCtx)
		resp, err := l.Insert(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
