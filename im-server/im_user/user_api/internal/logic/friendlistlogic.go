// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
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
	f := user_models.FriendModel{}
	friends := f.Friends(l.svcCtx.DB, 8)

	list := make([]types.UserInfoData, 0)

	for _, friend := range friends {
		item := types.UserInfoData{
			UserID: friend.SenderUserId,
		}
		list = append(list, item)
	}

	return &types.FriendListResponse{
		list,
		0,
	}, nil
}
