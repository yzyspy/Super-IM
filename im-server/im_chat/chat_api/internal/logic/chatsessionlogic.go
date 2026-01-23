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

type SessionResult struct {
	S_u uint
	R_u uint
}

// select * from (select least(sender_user_id, recv_user_id) as s_u,
//
// greatest(sender_user_id, recv_user_id) as r_u,
// count(id),
// max(created_at) as max_date,
// max(msg_preview)  as msg_preview,
// max(msg)  as msg
// from chat_models
// where sender_user_id = 8
// or recv_user_id = 8
// group by least(sender_user_id, recv_user_id),greatest(sender_user_id,recv_user_id))
//
// as subquery
// order by max_date
// limit 10 offset 0;
func (l *ChatSessionLogic) ChatSession(req *types.ChatSessionRequest) (resp *ChatSessionResponseList, err error) {
	db := l.svcCtx.DB
	userID := req.UserId

	// *“最近聊天列表”**查询逻辑。它的核心目的是将“我发给 A”和“A 发给我”的记录合并为同一条会话，并按最后一条消息的时间排序。
	subQuery := db.Table("chat_models").
		Select("LEAST(sender_user_id, recv_user_id) AS s_u, "+
			"GREATEST(sender_user_id, recv_user_id) AS r_u, "+
			"COUNT(id) AS count, "+
			"MAX(created_at) AS max_date, "+
			"max(msg)  as msg, "+
			"MAX(msg_preview)  as msg_preview").
		Where("sender_user_id = ? OR recv_user_id = ?", userID, userID).
		Group("LEAST(sender_user_id, recv_user_id), GREATEST(sender_user_id, recv_user_id)")

	session := make([]*SessionResult, 0)
	// 执行完整查询
	errdb := db.Table("(?) as subquery", subQuery).
		Order("max_date DESC"). // 通常最近聊天应是倒序
		Limit(10).
		Offset(0).
		Find(&session).Error

	list := make([]*types.ChatSessionResponse, 0)
	if len(session) > 0 {
		for _, s := range session {
			//senderUid := userID
			//if s.S_u == userID {
			//	senderUid = s.S_u
			//}
			list = append(list, &types.ChatSessionResponse{
				UserID: s.R_u,
			})
		}
		resp = &ChatSessionResponseList{
			List:  list,
			Count: len(list),
		}
	}
	if errdb != nil {
		err = errdb
	}
	return
}
