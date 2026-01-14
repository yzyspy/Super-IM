// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_chat/chat_api/internal/logic"
	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"
)

func chatSessionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatSessionRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}

		l := logic.NewChatSessionLogic(r.Context(), svcCtx)
		resp, err := l.ChatSession(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
