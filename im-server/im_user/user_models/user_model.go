package user_models

import "im-server/common/models"

type UserModel struct {
	models.Model
	Pwd      string `gorm:"column:password;size:255" json:"password"`
	Nickname string `gorm:"column:nickname;size:255" json:"nickname"`
	Abstract string `gorm:"column:abstract;size:255" json:"abstract"` //用户简介
	Avatar   string `gorm:"column:avatar;size:255" json:"avatar"`     //用户头像
	IP       string `gorm:"column:ip;size:255" json:"ip"`             //用户登录IP
	Addr     string `gorm:"column:addr;size:255" json:"addr"`         //用户登录地址
}
