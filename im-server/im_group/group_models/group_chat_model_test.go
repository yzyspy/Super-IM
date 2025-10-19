package group_models

import (
	"fmt"
	"im-server/common/models/ctype"
	"testing"
)

func TestGroupChatModel(t *testing.T) {
	fmt.Println("GroupChatModel111111111")
	t.Log("xxxx")

	g := GroupChatModel{
		SenderUserId: 1,
		GroupID:      101,
		MsgType:      1,
		MsgPreview:   "hello",
		Msg:          ctype.Msg{MsgType: 1, Content: nil},
		SystemMsg:    nil,
	}
	jsonStr, err := g.toJson()
	if err != nil {
		t.Errorf("toJson() error = %v", err)
		return
	}
	fmt.Println(jsonStr)

	var test GroupChatModel
	test.parseObj([]byte(jsonStr))
	fmt.Println("groupID:" + fmt.Sprintf("%d", test.GroupID))
	fmt.Println("")
}
