// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/im_user/user_rpc/types/user_rpc"

	"im-server/im_auth/auth_api/internal/svc"
	"im-server/im_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.Response, err error) {
	if (req.UserName == "") || (req.Password == "") {
		return &types.Response{
			Code: 400,
			Msg:  "user_name or password is empty",
		}, nil
	}
	//调用用户服务的rpc进行新用户的创建
	l.svcCtx.UserRpc.CreateUser(l.ctx, &user_rpc.UserCreateRequest{NickName: req.UserName, Password: req.Password})
	return &types.Response{
		Msg: "test",
	}, nil
	return
}
