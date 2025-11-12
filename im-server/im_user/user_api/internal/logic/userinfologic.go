// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"errors"
	"im-server/im_user/user_models"

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

func (l *UserInfoLogic) UserInfo(req *types.UserInfoRequest) (resp *types.UserInfoResponse, err error) {
	var user user_models.UserModel
	// gorm
	err = l.svcCtx.DB.Preload("UserConfModel").Take(&user, req.UserID).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	//fmt.Printf("query user model %+v\n", user)
	//fmt.Printf("query user model conf %s\n", *user.UserConfModel.RecallMessage)
	//byteData, _ := json.Marshal(user)
	//return &types.UserInfoResponse{Data: string(byteData)}

	return &types.UserInfoResponse{
		UserID:   user.UserConfModel.UserID,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}, nil
}
