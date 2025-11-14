// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_file/file_api/internal/logic"
	"im-server/im_file/file_api/internal/svc"
	"im-server/im_file/file_api/internal/types"
)

func ImageShowHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageShowRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}

		l := logic.NewImageShowLogic(r.Context(), svcCtx)
		err := l.ImageShow(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.Ok(w)
		}
	}
}
