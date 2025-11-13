package logic

import (
	"context"
	"github.com/pkg/errors"
	"im-server/im_user/user_models"

	"im-server/im_user/user_rpc/internal/svc"
	"im-server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(req *user_rpc.GetUserRequest) (*user_rpc.GetUserResponse, error) {
	var user user_models.UserModel
	// gorm
	err := l.svcCtx.DB.Preload("UserConfModel").Take(&user, req.UserId).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	//fmt.Printf("query user model %+v\n", user)
	//fmt.Printf("query user model conf %s\n", *user.UserConfModel.RecallMessage)
	//byteData, _ := json.Marshal(user)
	//return &types.UserInfoResponse{Data: string(byteData)}

	return &user_rpc.GetUserResponse{
		NickName: user.Nickname,
		Avator:   user.Avatar,
	}, nil
}
