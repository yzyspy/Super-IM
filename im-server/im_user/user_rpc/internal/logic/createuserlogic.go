package logic

import (
	"context"
	"fmt"

	"im-server/im_user/user_rpc/internal/svc"
	"im-server/im_user/user_rpc/types/user_rpc"

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

func (l *CreateUserLogic) CreateUser(in *user_rpc.UserCreateRequest) (*user_rpc.UserCreateResponse, error) {
	// todo: add your logic here and delete this line
	fmt.Println("CreateUserLogic yangzhongyu")
	return &user_rpc.UserCreateResponse{}, nil
}
