package group_models

import "../../common/models"

type GroupMemberModel struct {
	models.Model
	GroupID         uint   `gorm:"column:group_id" json:"group_id"`                 // 群ID
	UserID          uint   `gorm:"column:user_id" json:"user_id"`                   // 用户ID
	MemberNickName  string `gorm:"column:member_nick_name" json:"member_nick_name"` // 群成员昵称
	Role            int    `gorm:"column:role" json:"role"`                         // 群成员角色 3:普通成员 2:管理员 1:群主
	ProhibitionTime int    `gorm:"column:prohibition_time" json:"prohibition_time"` // 禁言时间
}
