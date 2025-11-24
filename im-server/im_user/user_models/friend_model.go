package user_models

import (
	"fmt"
	"gorm.io/gorm"
	"im-server/common/models"
)

// 好友表
type FriendModel struct {
	models.Model
	SenderUserId uint   `gorm:"column:sender_user_id" json:"sender_user_id"` //好友申请uid
	RecvUserId   uint   `gorm:"column:recv_user_id" json:"recv_user_id"`     //接受好友uid
	Notice       string `gorm:"column:notice;size:255" json:"notice"`        //好友备注【这个功能没有做，就显示对方昵称吧】

	SendUserModel *UserModel `gorm:"foreignkey:SenderUserId" json:"sender_user"`
	RecvUserModel *UserModel `gorm:"foreignkey:RecvUserId" json:"recv_user"`
}

// 判断A 和B是否是好友关系
func (f *FriendModel) IsFriend(db *gorm.DB, A uint, B uint) bool {
	err := db.Take(&f, "(send_user_id = ? and recv_user_id = ?) or (send_user_id = ? and recv_user_id = ?)", A, B, B, A).Error
	if err == nil {
		return true
	}
	fmt.Println(err) //record not found

	return false
}

// 查询A的全部好友
func (f *FriendModel) Friends(db *gorm.DB, A uint) (list []FriendModel) { //golang特性，命名返回值
	db.Preload("SendUserModel").Preload("RecvUserModel").Find(&list, "sender_user_id = ? or recv_user_id = ?", A, A)
	return
}
