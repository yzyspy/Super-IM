// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"
	"im-server/im_user/user_models"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendListLogic) FriendList(req *types.FriendListRequest) (resp *types.FriendListResponse, err error) {
	currentUserID := req.UserID

	f := user_models.FriendModel{}

	friends := f.Friends(l.svcCtx.DB, 8)

	list := make([]types.UserInfoData, 0)

	for _, friend := range friends {
		fmt.Printf("currentUserID: %d friend.RecvUserModel: %v\n", currentUserID, friend.RecvUserModel)
		if friend.SenderUserId == currentUserID { // 如果当前用户是好友的发起者，则查询接收者的信息
			item := types.UserInfoData{
				UserID:   friend.RecvUserId,
				Nickname: friend.RecvUserModel.Nickname,
				Avatar:   friend.RecvUserModel.Avatar,
				Abstract: friend.RecvUserModel.Abstract,
			}
			list = append(list, item)
		} else if friend.RecvUserId == currentUserID { // 如果当前用户是好友的结束者，则查询发送者的信息
			item := types.UserInfoData{
				UserID:   friend.SenderUserId,
				Nickname: friend.SendUserModel.Nickname,
				Avatar:   friend.SendUserModel.Avatar,
				Abstract: friend.SendUserModel.Abstract,
			}
			list = append(list, item)
		}
	}

	return &types.FriendListResponse{
		list,
		0,
	}, nil
}
