package logic

import (
	"context"
	"fmt"
	"im-server/im_user/user_models"

	"im-server/im_user/user_rpc/internal/svc"
	"im-server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserBatchLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserBatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserBatchLogic {
	return &GetUserBatchLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserBatchLogic) GetUserBatch(in *user_rpc.GetUserBatchRequest) (resp *user_rpc.GetUserBatchResponse, err error) {
	var userList []user_models.UserModel
	db_err := l.svcCtx.DB.Preload("UserConfModel").Where("id in (?)", in.UserIds).Find(&userList).Error
	if db_err != nil {
		fmt.Printf("GetUserBatch err: %v", db_err)
		return
	}
	if len(userList) == 0 {
		return
	}
	resp = &user_rpc.GetUserBatchResponse{}
	resp.Users = make(map[uint64]*user_rpc.GetUserResponse, 0)
	for _, user := range userList {
		resp.Users[uint64(user.ID)] = &user_rpc.GetUserResponse{
			NickName: user.Nickname,
			Avator:   user.Avatar,
			Abstract: user.Abstract,
		}
	}
	return
}
