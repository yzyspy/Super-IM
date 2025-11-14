// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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

		file, fileHead, read_file_err := r.FormFile("image")
		if read_file_err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: read_file_err.Error()})
			return
		}
		byteData, _ := io.ReadAll(file)

		imageType := r.FormValue("imageType")

		fileName := fileHead.Filename
		filePath := path.Join("uploads", imageType, fileName)
		fmt.Println(filePath)
		err := os.WriteFile(filePath, byteData, 0666)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
			return
		}

		l := logic.NewImageLogic(r.Context(), svcCtx)
		resp, err := l.Image(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.Response{Code: 400, Msg: err.Error()})
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
