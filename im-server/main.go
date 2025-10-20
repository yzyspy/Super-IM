package main

import (
	"im-server/im_chat/chat_models"
	"im-server/im_group/group_models"
	"im-server/im_user/user_models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:a12345678@tcp(127.0.0.1:3306)/my_qq?charset=utf8mb4&parseTime=True&loc=Local"
	db, error := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//初始化sys_user表
	db.AutoMigrate(
		&user_models.UserModel{},         //用户表
		&user_models.FriendModel{},       //好友表
		&user_models.FriendVerifyModel{}, //好友验证表
		&user_models.UserConfModel{},     // 聊天记录表

		&chat_models.ChatModel{}, // 聊天记录表

		&group_models.GroupChatModel{},   //群聊消息
		&group_models.GroupMemberModel{}, //群聊成员表
		&group_models.GroupModel{},       //群聊信息表
		&group_models.GroupVerifyModel{}, //加群申请表
	)

	if error != nil {
		panic("failed to connect database")
	}
	DB = db
}

func main() {
	// NewDB()
}
