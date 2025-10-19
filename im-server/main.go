package main

import (
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
		&user_models.UserModel{},
	)
	if error != nil {
		panic("failed to connect database")
	}
	DB = db
}

func main() {
	NewDB()
}
