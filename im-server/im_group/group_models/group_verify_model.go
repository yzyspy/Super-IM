package group_models

import "im-server/common/models/ctype"
import "im-server/common/models"

// 加群申请
type GroupVerifyModel struct {
	models.Model
	GroupID              uint                        `gorm:"column:group_id" json:"group_id"`                              // 群ID
	UserID               uint                        `gorm:"column:user_id" json:"user_id"`                                // 用户ID
	Status               uint8                       `gorm:"column:status" json:"status"`                                  //1:同意 2:拒绝 3.忽略
	AddtionalMessage     string                      `gorm:"column:additional_message;size:255" json:"additional_message"` //额外消息
	VerificationQuestion *ctype.VerificationQuestion `gorm:"column:verification_question" json:"verification_question"`    //验证问题
	Type                 uint8                       `gorm:"column:type" json:"type"`                                      //1:加群 2:退群
}
