// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGroupSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupSessionLogic {
	return &GroupSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupSessionLogic) GroupSession(req *types.GroupSessionRequest) (resp *types.GroupSessionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
