// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"im-server/common/models"
	"im-server/common/models/list_query"
	"im-server/im_chat/chat_models"

	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"

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

func (l *ChatHistoryLogic) ChatHistory(req *types.ChatHistoryRequest) (resp *types.ChatHistoryResponse, err error) {
	list_query.ListQuery(l.svcCtx.DB, chat_models.ChatModel{}, list_query.Option{
		PageInfo: models.PageInfo{
			Page:  req.Page,
			Limit: req.Limit,
		},
		Where: l.svcCtx.DB.Where("sender_user_id = ? or recv_user_id = ?", req.UserId, req.UserId),
	})
	//根据uid去重
	//批量uid查询用户的详情
	return
}
