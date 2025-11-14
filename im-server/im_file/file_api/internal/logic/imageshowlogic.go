// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"im-server/im_file/file_api/internal/svc"
	"im-server/im_file/file_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ImageShowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImageShowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImageShowLogic {
	return &ImageShowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ImageShowLogic) ImageShow(req *types.ImageShowRequest) error {

	return nil
}
