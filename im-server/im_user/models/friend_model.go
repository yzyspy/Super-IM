package models

import "../../common/models"

// 好友表
type FriendModel struct {
	models.Model
	SenderUserId uint   `gorm:"column:user_id" json:"user_id"`
	RecvUserId   uint   `gorm:"column:recv_user_id" json:"recv_user_id"`
	Notice       string `gorm:"column:notice" json:"notice"` //好友申请备注
}
