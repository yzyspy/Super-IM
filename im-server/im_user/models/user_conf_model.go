package models

import "../../common/models"

// 用户配置表
type UserConfModel struct {
	models.Model
	UserID               uint                  `gorm:"column:user_id" json:"user_id"`
	RecallMessage        *string               `gorm:"column:recall_message" json:"recall_message"`               //撤回消息的提示内容
	FriendOnline         bool                  `gorm:"column:friend_online" json:"friend_online"`                 //好友上线提醒
	Sound                bool                  `gorm:"column:sound" json:"sound"`                                 //收到消息是否播放声音
	SecureLink           bool                  `gorm:"column:secure_link" json:"secure_link"`                     //是否开启安全链接
	SavePwd              bool                  `gorm:"column:save_pwd" json:"save_pwd"`                           //是否保存密码
	SearchUser           uint8                 `gorm:"column:search_user" json:"search_user"`                     //别人查找你的方式 0:不允许查找 1:允许通过用户ID查询 2:允许通过昵称查询
	FriendVerification   uint8                 `gorm:"column:friend_verification" json:"friend_verification"`     //好友验证方式 0:不允许任何人添加 1:允许任何人添加 2：需要验证消息 3:需要回答问题 4:需要正确回答问题
	VerificationQuestion *VerificationQuestion `gorm:"column:verification_question" json:"verification_question"` //验证问题
	Online               bool                  `gorm:"column:online" json:"online"`                               //是否在线
}

type VerificationQuestion struct {
	Question1 *string `gorm:"column:question" json:"question1"` //验证问题
	Answer1   *string `gorm:"column:answer" json:"answer1"`     //验证问题的答案
	Question2 *string `gorm:"column:question" json:"question2"` //验证问题
	Answer2   *string `gorm:"column:answer" json:"answer2"`     //验证问题的答案
	Question3 *string `gorm:"column:question" json:"question3"` //验证问题
	Answer3   *string `gorm:"column:answer" json:"answer3"`     //验证问题的答案
}
