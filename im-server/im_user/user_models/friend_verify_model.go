package user_models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

// 好友申请
type FriendVerifyModel struct {
	models.Model
	SenderUserId         uint                        `gorm:"column:user_id" json:"user_id"`
	RecvUserId           uint                        `gorm:"column:recv_user_id" json:"recv_user_id"`
	Notice               string                      `gorm:"column:notice" json:"notice"`                               //好友申请备注
	Status               uint8                       `gorm:"column:status" json:"status"`                               //1:同意 2:拒绝 3.忽略
	AddtionalMessage     string                      `gorm:"column:additional_message" json:"additional_message"`       //额外消息
	VerificationQuestion *ctype.VerificationQuestion `gorm:"column:verification_question" json:"verification_question"` //验证问题
}
