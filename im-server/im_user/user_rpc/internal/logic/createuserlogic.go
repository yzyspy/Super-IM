package logic

import (
	"context"
	"im-server/im_user/user_models"
	"im-server/im_user/user_rpc/internal/svc"
	"im-server/im_user/user_rpc/types/user_rpc"
	"im-server/utils/pwd"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(req *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	hashPwd := pwd.HashPwd(req.Password)
	user_model := user_models.UserModel{
		Pwd:      hashPwd,
		Nickname: req.NickName,
	}
	//保存用户基本信息
	l.svcCtx.DB.Save(&user_model)

	user_conf_model := user_models.UserConfModel{
		UserID:        user_model.ID,
		RecallMessage: nil,
		FriendOnline:  false,
		Sound:         false,
		SecureLink:    false,
		SavePwd:       false,
		SearchUser:    0,
	}
	//保存用户配置信息
	l.svcCtx.DB.Save(&user_conf_model)

	l.Logger.Infof("注册成功, user_id: %d", user_model.ID)
	return &user_rpc.UserCreateResponse{
		UserId: uint64(user_model.ID),
	}, nil
}
