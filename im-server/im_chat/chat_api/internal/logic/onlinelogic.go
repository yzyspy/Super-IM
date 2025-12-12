// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"fmt"
	"im-server/utils/jwt"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"im-server/im_chat/chat_api/internal/svc"
	"im-server/im_chat/chat_api/internal/types"
)

type OnLineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type UserInfo struct {
	UserId   int64  `json:"user_id"`
	NickName string `json:"nick_name"`
	Avatar   string `json:"avatar"`
}

type UserWsInfo struct {
	UserInfo UserInfo
	Conn     *websocket.Conn
}

// 保存所有的websocket连接，key为userId
var UserWsMap = map[int64]*UserWsInfo{}

func NewOnLineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OnLineLogic {
	return &OnLineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OnLineLogic) OnLine(req *types.OnLineRequest, w http.ResponseWriter, r *http.Request) (resp *types.OnLineResponse, err error) {
	payLoad, parseError := jwt.ParseJWT(req.Token)
	if parseError != nil {
		fmt.Println(parseError)
	}
	uid, _ := strconv.Atoi(payLoad.UserID)
	fmt.Printf("uid=%d webocket 上线了\n", uid)
	//http协议升级为websocket协议
	var upgrader = websocket.Upgrader{
		// 读缓冲区大小
		ReadBufferSize: 1024,
		// 写缓冲区大小
		WriteBufferSize: 1024,
		// 允许跨域（重要！生产环境应根据需要精确配置）
		CheckOrigin: func(r *http.Request) bool {
			// 允许所有请求来源，仅为演示。
			// 生产环境应严格检查 r.Header.Get("Origin")
			return true
		},
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Upgrade error:%v\n", err)
		return
	}

	defer func() {
		conn.Close() // 确保连接在函数退出时关闭
		delete(UserWsMap, int64(uid))
		fmt.Printf("uid=%d webocket 下线了\n", uid)
	}

	//l.svcCtx.UserRpc.
	UserWsMap[int64(uid)] = &UserWsInfo{
		UserInfo: UserInfo{
			UserId:   int64(uid),
			NickName: payLoad.NickName,
			Avatar:   payLoad.Avatar,
		},
		Conn: conn,
	}

	// 2. WebSocket 连接建立成功，开始处理消息
	fmt.Println("WebSocket client connected.")

	// 消息循环
	for {
		// 读取消息
		// messageType (int): 消息类型 (Text=1, Binary=2)
		// message ([]byte): 消息内容
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			// 客户端断开连接或读取错误
			fmt.Printf("Read error:", err)
			break
		}

		fmt.Printf("Received: %s (Type: %d)", message, messageType)

		// 写入消息 (回复给客户端)
		response := []byte("Server received your message: " + string(message))
		if err := conn.WriteMessage(websocket.TextMessage, response); err != nil {
			fmt.Println("Write error:", err)
			break
		}
	}

	resp = &types.OnLineResponse{
		Status: true,
	}
	return
}
