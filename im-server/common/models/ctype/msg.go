package ctype

type SystemMsg struct {
	Type uint8 `gorm:"column:msg_type" json:"msg_type"` //违规类型
}

type Msg struct {
	MsgType      uint8         `gorm:"column:msg_type" json:"msg_type"`             //消息类型 1:文本 2:图片
	Content      *string       `gorm:"column:content" json:"content"`               //消息内容
	ImageMsg     *ImageMsg     `gorm:"column:image_msg" json:"image_msg"`           //图片消息
	VideoMsg     *VideoMsg     `gorm:"column:video_msg" json:"video_msg"`           //视频消息
	FileMsg      *FileMsg      `gorm:"column:file_msg" json:"file_msg"`             //文件消息
	VoiceMsg     *VoiceMsg     `gorm:"column:voice_msg" json:"voice_msg"`           //语音消息
	VideoCallMsg *VideoCallMsg `gorm:"column:video_call_msg" json:"video_call_msg"` //视频通话消息
	VoiceCallMsg *VoiceCallMsg `gorm:"column:voice_call_msg" json:"voice_call_msg"` //语音通话消息
	WitchdrawMsg *WitchdrawMsg `gorm:"column:witchdraw_msg" json:"witchdraw_msg"`   //撤回消息
	ReplyMsg     *ReplyMsg     `gorm:"column:reply_msg" json:"reply_msg"`           //回复消息
	QuoteMsg     *QuoteMsg     `gorm:"column:quote_msg" json:"quote_msg"`           //引用消息
	AtMsg        *AtMsg        `gorm:"column:at_msg" json:"at_msg"`                 //@消息
}

type ImageMsg struct {
	Title string `gorm:"column:title" json:"title"` //图片标题
	Src   string `gorm:"column:src" json:"src"`     //图片地址
}
type VideoMsg struct {
	Title string `gorm:"column:title" json:"title"` //图片标题
	Src   string `gorm:"column:src" json:"src"`     //图片地址
	Time  int    `gorm:"column:time" json:"time"`   //视频时长
}
type FileMsg struct {
	Title string `gorm:"column:title" json:"title"` //图片标题
	Src   string `gorm:"column:src" json:"src"`     //图片地址
	Size  int    `gorm:"column:size" json:"size"`   //文件大小
	Type  string `gorm:"column:type" json:"type"`   //文件类型
}
type VoiceMsg struct {
	Title string `gorm:"column:title" json:"title"` //图片标题
	Src   string `gorm:"column:src" json:"src"`     //图片地址
}

type VoiceCallMsg struct {
	StartTime time.Time `gorm:"column:start_time" json:"start_time"` //通话开始时间
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`     //通话结束时间
	EndReason string    `gorm:"column:end_reason" json:"end_reason"` //通话结束原因
}

type VideoCallMsg struct {
	StartTime time.Time `gorm:"column:start_time" json:"start_time"` //通话开始时间
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`     //通话结束时间
	EndReason string    `gorm:"column:end_reason" json:"end_reason"` //通话结束原因
}

// 撤回消息
type WitchdrawMsg struct {
	OriginMsg *Msg    `gorm:"column:origin_msg" json:"origin_msg"` //原始消息
	Content   *string `gorm:"column:content" json:"content"`       //撤回消息提示词
}

// 回复消息
type ReplyMsg struct {
	MsgId   uint    `gorm:"column:msg_id" json:"msg_id"`   //引用消息ID
	Content *string `gorm:"column:content" json:"content"` //回复消息内容,目前只支持回复文本，不能回复语音
	Msg     *Msg    `gorm:"column:msg" json:"msg"`
}

// 引用消息
type QuoteMsg struct {
	MsgId   uint    `gorm:"column:msg_id" json:"msg_id"`   //引用消息ID
	Content *string `gorm:"column:content" json:"content"` //回复消息内容,目前只支持回复文本，不能回复语音
	Msg     *Msg    `gorm:"column:msg" json:"msg"`
}

type AtMsg struct {
	MsgId   uint    `gorm:"column:msg_id" json:"msg_id"`   //引用消息ID
	Content *string `gorm:"column:content" json:"content"` //回复消息内容,目前只支持回复文本，不能回复语音
	Msg     *Msg    `gorm:"column:msg" json:"msg"`
}
