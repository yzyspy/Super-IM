// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/im_auth/auth_models"
	"im-server/utils/jwt"
	"im-server/utils/pwd"
	"strconv"
	"time"

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

	l.svcCtx.Redis.Set("yzy", "123", time.Duration(time.Duration.Seconds(60)))

	user_model := auth_models.UserModel{}
	l.svcCtx.DB.Take(&user_model, "nickname = ?", req.UserName)
	//用户不存在
	if user_model.ID == 0 {
		return &types.LoginResponse{
			Code: 1,
			Msg:  "user not exist, check your username",
			Data: types.LoginInfo{},
		}, nil
	}
	// 判断密码是否正确
	is_pwd_right := pwd.CheckPwd(user_model.Pwd, req.Password)
	if !is_pwd_right {
		return &types.LoginResponse{
			Code: 1,
			Msg:  "password error, check your password",
			Data: types.LoginInfo{},
		}, nil
	}
	// 生成jwt token
	token, _ := jwt.GenerateJWT(strconv.Itoa(int(user_model.ID)), user_model.Nickname)

	return &types.LoginResponse{
		Code: 0,
		Msg:  "success",
		Data: types.LoginInfo{
			Token: token,
		},
	}, nil
}
