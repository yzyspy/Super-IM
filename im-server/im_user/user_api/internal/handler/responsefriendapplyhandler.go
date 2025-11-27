// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_user/user_api/internal/logic"
	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"
)

func responseFriendApplyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResponseFriendApplyRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 401, Msg: err.Error()})
			return
		}

		l := logic.NewResponseFriendApplyLogic(r.Context(), svcCtx)
		resp, err := l.ResponseFriendApply(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 402, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
