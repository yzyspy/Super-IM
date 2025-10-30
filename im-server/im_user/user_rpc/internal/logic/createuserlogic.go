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
	l.svcCtx.DB.Save(&user_model)
	l.Logger.Infof("注册成功, user_id: %d", user_model.ID)
	return &user_rpc.UserCreateResponse{
		UserId: uint64(user_model.ID),
	}, nil
}
