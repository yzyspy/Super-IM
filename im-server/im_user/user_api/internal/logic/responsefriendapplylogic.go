// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/common/models"
	"im-server/im_user/user_models"

	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResponseFriendApplyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResponseFriendApplyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResponseFriendApplyLogic {
	return &ResponseFriendApplyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResponseFriendApplyLogic) ResponseFriendApply(req *types.ResponseFriendApplyRequest) (resp *types.ResponseFriendApplyResponse, err error) {
	model := user_models.FriendVerifyModel{
		Model: models.Model{
			ID: uint(req.Friend_verify_id), // 明确指定嵌套结构体的字段
		},
		Status: uint8(req.Status),
	}
	l.svcCtx.DB.Updates(&model)
	resp = &types.ResponseFriendApplyResponse{
		Code: 0,
		Msg:  "success",
	}
	return
}
