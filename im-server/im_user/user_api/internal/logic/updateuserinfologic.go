// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/common/models"
	"im-server/im_user/user_models"
	"strconv"

	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UserInfoUpdateRequest, uidStr string) (resp *types.UserInfoUpdateResponse, err error) {
	uid, _ := strconv.Atoi(uidStr)
	userModel := user_models.UserModel{
		Model: models.Model{
			ID: uint(uid), // 明确指定嵌套结构体的字段
		},
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	}
	l.svcCtx.DB.Updates(&userModel)
	return &types.UserInfoUpdateResponse{
		Code: 0,
		Msg:  "success",
		Data: true,
	}, nil
}
