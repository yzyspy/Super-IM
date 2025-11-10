// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_auth/auth_api/internal/logic"
	"im-server/im_auth/auth_api/internal/svc"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		token := r.Header.Get("token")
		uid := r.Header.Get("uid")
		clientIp := r.Header.Get("X-Forwarded-For")
		fmt.Printf("logoutHandler clientIp: %s, token: %s, uid: %s\n", clientIp, token, uid)
		l.Logger.Info("logoutHandler2 clientIp: %s, token: %s, uid: %s", clientIp, token, uid)
		resp, err := l.Logout(token)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
