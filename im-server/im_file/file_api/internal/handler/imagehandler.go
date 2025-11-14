// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"io"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"im-server/im_file/file_api/internal/logic"
	"im-server/im_file/file_api/internal/svc"
	"im-server/im_file/file_api/internal/types"
)

func ImageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ImageRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}

		file, _, read_file_err := r.FormFile("image")
		if read_file_err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: read_file_err.Error()})
			return
		}
		io.ReadAll(file)

		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
