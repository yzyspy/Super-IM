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
	//更新好友申请的状态
	l.svcCtx.DB.Updates(&user_models.FriendVerifyModel{
		Model: models.Model{
			ID: uint(req.Friend_verify_id), // 明确指定嵌套结构体的字段
		},
		Status: uint8(req.Status), // 同意为1，拒绝为2
	})

	if req.Status == 1 {
		// 同意添加好友
		friend_verify_model := user_models.FriendVerifyModel{
			Model: models.Model{
				ID: uint(req.Friend_verify_id), // 明确指定嵌套结构体的字段
			},
		}
		l.svcCtx.DB.Take(&friend_verify_model)

		friend_model := user_models.FriendModel{
			SenderUserId: friend_verify_model.SenderUserId,
			RecvUserId:   friend_verify_model.RecvUserId,
		}
		l.svcCtx.DB.Create(&friend_model)
	}
	resp = &types.ResponseFriendApplyResponse{
		Code: 0,
		Msg:  "success",
		Data: true,
	}
	return
}
