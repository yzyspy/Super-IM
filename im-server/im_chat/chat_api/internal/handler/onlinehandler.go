// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"github.com/gorilla/websocket"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_chat/chat_api/internal/logic"
	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"
)

type UserInfo struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

type UserWsInfo struct {
	UserInfo UserInfo
	Conn     *websocket.Conn
}

// 保存所有的websocket连接，key为userId
var UserWsMap = map[int64]*UserWsInfo{}

func onLineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OnLineRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}

		l := logic.NewOnLineLogic(r.Context(), svcCtx)
		resp, err := l.OnLine(&req, w, r)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
