package chat_models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

// 聊天记录表
type ChatModel struct {
	models.Model
	SenderUserId uint             `gorm:"column:sender_user_id" json:"sender_user_id"`    //发送用户ID
	RecvUserId   uint             `gorm:"column:recv_user_id" json:"recv_user_id"`        //接收用户ID
	MsgType      uint8            `gorm:"column:msg_type" json:"msg_type"`                //消息类型 1:文本 2:图片
	MsgPreview   string           `gorm:"column:msg_preview;size:255" json:"msg_preview"` //消息预览
	Msg          ctype.Msg        `gorm:"column:msg" json:"msg"`                          //消息内容
	SystemMsg    *ctype.SystemMsg `gorm:"column:system_msg" json:"system_msg"`            //系统消息
}
