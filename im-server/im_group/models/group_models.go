package models

import "../../common/models"
import "../../common/models/ctype"

type GroupModel struct {
	models.Model
	Title                string                      `gorm:"column:title" json:"title"`       //群名称
	Abstract             string                      `gorm:"column:abstract" json:"abstract"` //群描述
	Avatar               string                      `gorm:"column:avatar" json:"avatar"`     //群头像
	Creator              uint                        `gorm:"column:creator" json:"creator"`   //群创建者
	IsSearchable         bool                        `gorm:"column:is_searchable" json:"is_searchable"`
	Verification         uint                        `gorm:"column:verification" json:"verification"`                   // 好友验证方式 0:不允许任何人添加 1:允许任何人添加 2：需要验证消息 3:需要回答问题 4:需要正确回答问题
	VerificationQuestion *ctype.VerificationQuestion `gorm:"column:verification_question" json:"verification_question"` // 验证问题
	IsInvite             bool                        `gorm:"column:is_invite" json:"is_invite"`                         // 是否允许邀请好友加入
	IsTemporarySession   bool                        `gorm:"column:is_temporary_session" json:"is_temporary_session"`   // 是否临时会话
	IsProhibition        bool                        `gorm:"column:is_prohibition" json:"is_prohibition"`               // 是否开启全员禁言
}
