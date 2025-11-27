// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"im-server/im_user/user_models"
	"strconv"

	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyFriendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyFriendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyFriendLogic {
	return &ApplyFriendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ApplyFriendLogic) ApplyFriend(req *types.ApplyFriendRequest, currentUserIDStr string) (resp *types.ApplyFriendResponse, err error) {
	// 检查用户是否存在
	var user user_models.UserModel
	// gorm
	find_user_err := l.svcCtx.DB.Preload("UserConfModel").Take(&user, req.UserID).Error

	if find_user_err != nil {
		return nil, errors.New("用户不存在")
	}
	//校验是否已经是好友
	currentUid, atoi_err := strconv.Atoi(currentUserIDStr)
	if atoi_err != nil {
		return nil, errors.New("current user id error")
	}
	f := user_models.FriendModel{}
	isFriend := f.IsFriend(l.svcCtx.DB, req.UserID, uint(currentUid))
	if isFriend {
		return nil, errors.New("已经是好友了")
	}
	// insert 好友验证表,等待被邀请方同意
	model := user_models.FriendVerifyModel{
		SenderUserId: uint(currentUid),
		RecvUserId:   req.UserID,
		Status:       0,
	}
	l.svcCtx.DB.Create(&model)

	resp = &types.ApplyFriendResponse{
		Code: 0,
		Msg:  "success",
		Data: true,
	}
	return
}
