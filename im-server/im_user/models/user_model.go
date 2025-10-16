package models

import "../../common/models"

type UserModel struct {
	models.Model
	Pwd      string `gorm:"column:password" json:"password"`
	Nickname string `gorm:"column:nickname" json:"nickname"`
	Abstract string `gorm:"column:abstract" json:"abstract"` //用户简介
	Avatar   string `gorm:"column:avatar" json:"avatar"`     //用户头像
}
