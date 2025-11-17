// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"

	"im-server/im_file/file_api/internal/svc"
	"im-server/im_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageLogic {
	return &ImageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageLogic) Image(req *types.ImageRequest, filePath string) (resp *types.ImageResponsse, err error) {
	filePathUrl := fmt.Sprintf("%s%s", "/api/file/", filePath)
	return &types.ImageResponsse{
		Url: filePathUrl,
	}, nil
}
