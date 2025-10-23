package group_models

import (
	"im-server/common/models"
	"im-server/common/models/ctype"
)

// 群聊消息
type GroupChatModel struct {
	models.Model
	SenderUserId uint             `gorm:"column:sender_user_id" json:"sender_user_id"`    //发送用户ID
	GroupID      uint             `gorm:"column:group_id" json:"group_id"`                //群ID
	MsgType      uint8            `gorm:"column:msg_type" json:"msg_type"`                //消息类型 1:文本 2:图片
	MsgPreview   string           `gorm:"column:msg_preview;size:255" json:"msg_preview"` //消息预览
	Msg          ctype.Msg        `gorm:"column:msg" json:"msg"`                          //消息内容
	SystemMsg    *ctype.SystemMsg `gorm:"column:system_msg" json:"system_msg"`            //系统消息
}

//func (g GroupChatModel) toJson() (string, error) {
//	b, err := json.Marshal(g)
//	return string(b), err
//}
//
//// 这个地方要用指针 接受者
//func (g *GroupChatModel) parseObj(obj any) error {
//	return json.Unmarshal(obj.([]byte), g)
//}
//
//// 这个地方要用指针 接受者
//func (g *GroupChatModel) parseObj2(jsonStr string) error {
//	return json.Unmarshal([]byte(jsonStr), g)
//}
