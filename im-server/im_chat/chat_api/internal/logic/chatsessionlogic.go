// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/im_user/user_rpc/types/user_rpc"

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
	Least_uid    uint64
	Greatest_uid uint64
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
	userID := uint64(req.UserId)

	// *“最近聊天列表”**查询逻辑。它的核心目的是将“我发给 A”和“A 发给我”的记录合并为同一条会话，并按最后一条消息的时间排序。
	// 还需要展示两个人对话的最后一条消息的预览，这里的MAX(msg_preview)有问题，应该是最大的created_at的那个消息
	// select msg from chat_models where ((sender_user_id = ? OR recv_user_id = ?) or (sender_user_id = ? OR recv_user_id = ?))
	// order by created_at limit 1
	subQuery := db.Table("chat_models").
		Select("LEAST(sender_user_id, recv_user_id) AS least_uid, "+
			"GREATEST(sender_user_id, recv_user_id) AS greatest_uid, "+
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

			oppositeUid := uint64(0)
			if s.Greatest_uid == userID {
				oppositeUid = s.Least_uid
			} else if s.Least_uid == userID {
				oppositeUid = s.Greatest_uid
			}
			oppositeUids := make([]uint64, 0)
			oppositeUids = append(oppositeUids, oppositeUid)
			getUserBatchRequest := &user_rpc.GetUserBatchRequest{
				UserIds: oppositeUids,
			}
			userInfos, _ := l.svcCtx.UserRpc.GetUserBatch(l.ctx, getUserBatchRequest)
			list = append(list, &types.ChatSessionResponse{
				UserID:   uint(oppositeUid),
				Nickname: userInfos.Users[oppositeUid].NickName,
				Avatar:   userInfos.Users[oppositeUid].Avator,
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
