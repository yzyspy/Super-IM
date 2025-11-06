// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_auth/auth_api/internal/logic"
	"im-server/im_auth/auth_api/internal/svc"
	"im-server/im_auth/auth_api/internal/types"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AuthenticationRequest
		if err := httpx.Parse(r, &req); err != nil {
			//httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}
		//token := r.Header.Get("token")
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		clientIp := r.Header.Get("X-Forwarded-For")
		logx.Info("authenticationHandler clientIp: %s", clientIp)
		resp, err := l.Authentication(&req)
		if err != nil {
			//	httpx.ErrorCtx(r.Context(), w, err)
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
