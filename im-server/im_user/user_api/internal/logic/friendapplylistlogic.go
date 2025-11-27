// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendApplyListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFriendApplyListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendApplyListLogic {
	return &FriendApplyListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendApplyListLogic) FriendApplyList(req *types.FriendApplyListRequest, currentUserId int) (resp *types.FriendApplyListResponse, err error) {
	// 查询好友申请表，所有recv_uid_id == currentUserId的记录
	return
}
