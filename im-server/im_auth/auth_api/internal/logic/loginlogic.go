// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/im_auth/auth_models"

	"im-server/im_auth/auth_api/internal/svc"
	"im-server/im_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	user_model := auth_models.UserModel{}
	l.svcCtx.DB.Take(&user_model, "nickname = ?", req.UserName)

	fmt.Printf("get user_model: %v\n", user_model)
	// 校验用户密码是否正确

	// 生成jwt token

	return &types.LoginResponse{
		Code: 0,
		Msg:  "success",
		Data: types.LoginInfo{
			Token: "123456",
		},
	}, nil
}
