// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/utils/jwt"
	"strconv"
	"strings"

	"im-server/im_auth/auth_api/internal/svc"
	"im-server/im_auth/auth_api/internal/types"

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

func isWhiteList(requestUrl string) bool {
	whiteList := []string{"/login", "/register"}
	for _, url := range whiteList {
		if strings.Contains(requestUrl, url) {
			return true
		}
	}
	return false
}

func (l *AuthenticationLogic) Authentication(req *types.AuthenticationRequest) (resp *types.AuthenticationResponse, err error) {
	requestUrl := req.ValidPath
	fmt.Printf("authenticationHandler requestUrl %s\n", requestUrl)
	fmt.Printf("authenticationHandler token %s\n", req.Token)
	//判断请求url是否在认证白名单中，如果在白名单中，直接返回认证成功，不校验token
	//login 、logout、authentication 这些认证接口是不需要登录的，直接返回认证成功
	if isWhiteList(requestUrl) {
		logx.Info("请求url在认证白名单中，直接返回认证成功", requestUrl)
		return &types.AuthenticationResponse{
			Code: 0,
			Msg:  "认证成功",
		}, nil
	}
	token := req.Token
	if token == "" {
		return &types.AuthenticationResponse{
			Code: 401,
			Msg:  "token is required",
		}, nil
	}
	//校验token是否已经过期
	payLoad, parseError := jwt.ParseJWT(token)
	if parseError != nil {
		return &types.AuthenticationResponse{
			Code: 401,
			Msg:  "认证失败",
		}, nil
	}
	//用户主动退出登录了
	key := fmt.Sprintf("logout_%s", token)
	val, err := l.svcCtx.Redis.Get(key).Result()
	fmt.Printf("authenticationHandler redis val %s\n", val)
	if err == nil && val == "1" {
		return &types.AuthenticationResponse{
			Code: 402,
			Msg:  "用户已退出登录",
		}, nil
	}
	uid, _ := strconv.Atoi(payLoad.UserID)
	return &types.AuthenticationResponse{
		Code:   0,
		Msg:    "认证成功",
		UserId: uint(uid),
		Role:   0,
	}, nil
}
