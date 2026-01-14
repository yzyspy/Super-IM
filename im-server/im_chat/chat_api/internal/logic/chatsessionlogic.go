// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"

	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatSessionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatSessionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatSessionLogic {
	return &ChatSessionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChatSessionResponseList struct {
	List  []*types.ChatSessionResponse `json:"list"`
	Count int                          `json:"count"`
}

// select * from (select least(sender_user_id, recv_user_id) as s_u,
//
//	greatest(sender_user_id, recv_user_id) as r_u,
//	count(id),
//	max(created_at),
//	max(msg_preview)  as max_date
//	from chat_models
//	where sender_user_id = 8
//	or recv_user_id = 8
//	group by least(sender_user_id, recv_user_id),greatest(sender_user_id,recv_user_id))
//
// as subquery
// order by max_date
// limit 10 offset 0;
func (l *ChatSessionLogic) ChatSession(req *types.ChatSessionRequest) (resp *types.ChatSessionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
