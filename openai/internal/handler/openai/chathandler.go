package openai

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-chatgpt/openai/internal/logic/openai"
	"go-zero-chatgpt/openai/internal/svc"
	"go-zero-chatgpt/openai/internal/types"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := openai.NewChatLogic(r.Context(), svcCtx)
		resp, err := l.Chat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
