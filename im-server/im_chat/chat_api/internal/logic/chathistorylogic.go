// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/common/models"
	"im-server/common/models/ctype"
	"im-server/common/models/list_query"
	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"
	"im-server/im_chat/chat_models"
	"im-server/im_user/user_rpc/types/user_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatHistoryLogic {
	return &ChatHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type ChatHistory struct {
	ID        uint             `json:"id"`
	UserID    uint             `json:"user_id"`
	Avatar    string           `json:"avatar"`
	Nickname  string           `json:"nickname"`
	CreatedAt string           `json:"created_at"`
	Msg       *ctype.Msg       `json:"msg"`
	SystemMsg *ctype.SystemMsg `json:"system_msg"`
	IsMe      bool             `json:"is_me"`
}

type ChatHistoryResponseList struct {
	List  []*ChatHistory `json:"list"`
	Count int            `json:"count"`
}

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *ChatHistoryResponseList, err error) {
	list, count, err := list_query.ListQuery(l.svcCtx.DB, chat_models.ChatModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},

		Where: l.svcCtx.DB.Where("(sender_user_id = ? and recv_user_id = ?) or (sender_user_id = ? and recv_user_id = ?)",
			req.FriendId, req.UserId, req.UserId, req.FriendId),
	})
	if count == 0 {
		return nil, fmt.Errorf("no chat history found")
	}

	uids := make([]uint64, 0)
	for _, item := range list {
		uids = append(uids, uint64(item.SenderUserId))
		uids = append(uids, uint64(item.RecvUserId))
	}
	//根据uid去重
	//批量uid查询用户的详情
	getUserBatchRequest := &user_rpc.GetUserBatchRequest{
		UserIds: uids,
	}
	userInfos, err := l.svcCtx.UserRpc.GetUserBatch(l.ctx, getUserBatchRequest)
	if err != nil {
		return
	}

	chatHistoryList := make([]*ChatHistory, 0)
	for _, item := range list {
		userInfo := userInfos.Users[uint64(item.SenderUserId)]
		if userInfo == nil {
			continue
		}

		chatHistory := &ChatHistory{
			ID:        item.ID,
			UserID:    item.SenderUserId,
			Avatar:    userInfo.Avator,
			Nickname:  userInfo.NickName,
			Msg:       &item.Msg,
			CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		}

		//if chatHistory.SendUser.ID == req.UserId {
		//	chatHistory.IsMe = true
		//}

		chatHistoryList = append(chatHistoryList, chatHistory)
	}
	resp = &ChatHistoryResponseList{
		List:  chatHistoryList,
		Count: int(count),
	}
	return
}
