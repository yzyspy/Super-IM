// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/utils/jwt"
	"time"

	"im-server/im_auth/auth_api/internal/svc"
	"im-server/im_auth/auth_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogoutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogoutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogoutLogic {
	return &LogoutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogoutLogic) Logout(token string) (resp *types.Response, err error) {
	if token == "" {
		return &types.Response{
			Code: 401,
			Msg:  "token is required",
		}, nil
	}
	payLoad, parseTokenError := jwt.ParseJWT(token)
	if parseTokenError != nil {
		return &types.Response{
			Code: 401,
			Msg:  "token invalid",
		}, nil
	}
	expireTime := payLoad.ExpiresAt - time.Now().Unix()
	key := fmt.Sprintf("logout_%s", token)
	l.svcCtx.Redis.SetNX(key, "1", time.Duration(expireTime)*time.Second)

	return &types.Response{
		Code: 0,
		Msg:  "logout success",
	}, nil
}
