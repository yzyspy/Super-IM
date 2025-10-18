package models

import "../../common/models"
import "../../common/models/ctype"

// 群聊消息
type GroupChatModel struct {
	models.Model
	SenderUserId uint             `gorm:"column:sender_user_id" json:"sender_user_id"` //发送用户ID
	GroupID      uint             `gorm:"column:group_id" json:"group_id"`             //群ID
	MsgType      uint8            `gorm:"column:msg_type" json:"msg_type"`             //消息类型 1:文本 2:图片
	MsgPreview   string           `gorm:"column:msg_preview" json:"msg_preview"`       //消息预览
	Msg          ctype.Msg        `gorm:"column:msg" json:"msg"`                       //消息内容
	SystemMsg    *ctype.SystemMsg `gorm:"column:system_msg" json:"system_msg"`         //系统消息
}
