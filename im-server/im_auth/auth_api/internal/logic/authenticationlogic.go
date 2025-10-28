// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/im_auth/auth_api/internal/svc"
	"im-server/im_auth/auth_api/internal/types"
	"im-server/utils/jwt"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthenticationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAuthenticationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthenticationLogic {
	return &AuthenticationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AuthenticationLogic) Authentication(token string) (resp *types.Response, err error) {
	if token == "" {
		return &types.Response{
			Code: 401,
			Msg:  "token is required",
		}, nil
	}
	//校验token是否已经过期
	payLoad, parseError := jwt.ParseJWT(token)
	if parseError != nil {
		return &types.Response{
			Code: 401,
			Msg:  "认证失败",
		}, nil
	}
	//用户主动退出登录了
	key := fmt.Sprintf("logout_%d", payLoad.UserID)
	val, err := l.svcCtx.Redis.Get(key).Result()
	if err == nil && val == "1" {
		return &types.Response{
			Code: 402,
			Msg:  "用户已退出登录",
		}, nil
	}

	return &types.Response{
		Code: 0,
		Msg:  "认证成功",
	}, nil
}
