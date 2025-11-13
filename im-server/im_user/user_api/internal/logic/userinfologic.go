// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/im_user/user_rpc/types/user_rpc"

	"im-server/im_user/user_api/internal/svc"
	"im-server/im_user/user_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 查询当前登录用户信息

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user_rpc.GetUserRequest{UserId: uint64(req.UserID)})
	return &types.UserInfoResponse{
		Nickname: user.NickName,
		Avatar:   user.Avator,
	}, nil
}
