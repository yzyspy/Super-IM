package user_models

import "im-server/common/models"

// 好友表
type FriendModel struct {
	models.Model
	SenderUserId uint   `gorm:"column:user_id" json:"user_id"`           //好友申请uid
	RecvUserId   uint   `gorm:"column:recv_user_id" json:"recv_user_id"` //接受好友uid
	Notice       string `gorm:"column:notice;size:255" json:"notice"`    //好友备注【这个功能没有做，就显示对方昵称吧】
}
