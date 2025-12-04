// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"
	"strconv"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_user/user_api/internal/logic"
	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"
)

func friendDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FriendDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}

		l := logic.NewFriendDeleteLogic(r.Context(), svcCtx)

		uid := r.Header.Get("uid")
		uidInt, _ := strconv.Atoi(uid)

		resp, err := l.FriendDelete(&req, uidInt)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
