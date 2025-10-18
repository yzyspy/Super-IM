package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"im-server/im_user/user_models"
)

var DB *gorm.DB

func NewDB() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/my_shopping_mall?charset=utf8mb4&parseTime=True&loc=Local"
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

}
