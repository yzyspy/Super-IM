// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/utils/jwt"
	"strconv"

	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OnLineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOnLineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnLineLogic {
	return &OnLineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnLineLogic) OnLine(req *types.OnLineRequest) (resp *types.OnLineResponse, err error) {
	payLoad, parseError := jwt.ParseJWT(req.Token)
	if parseError != nil {
		fmt.Println(parseError)
	}
	uid, _ := strconv.Atoi(payLoad.UserID)
	fmt.Printf("uid=%d webocket 上线了\n", uid)
	resp = &types.OnLineResponse{
		Status: true,
	}
	return
}
