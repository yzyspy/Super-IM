// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"im-server/common/models/ctype"
	"im-server/im_user/user_rpc/types/user_rpc"
	"im-server/utils/jwt"
	"net/http"
	"strconv"
	"time"

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
var UserWsMap = map[uint]*UserWsInfo{}

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
		conn.Close()
		delete(UserWsMap, uint(uid))
		fmt.Printf("uid=%d webocket 下线了\n", uid)
	}()

	// 1. 从数据库查询用户信息
	userResp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user_rpc.GetUserRequest{
		UserId: uint64(uid),
	})
	if err != nil {
		fmt.Printf("GetUser error:%v\n", err)
		return
	}
	//保存当前用户的长链接
	UserWsMap[uint(uid)] = &UserWsInfo{
		UserInfo: UserInfo{
			UserId:   int64(uid),
			NickName: userResp.NickName,
			Avatar:   userResp.Avator,
		},
		Conn: conn,
	}
	// 2. WebSocket 连接建立成功，开始处理消息
	fmt.Printf("WebSocket client connected. uid=%s nick_name=%s", int64(uid), userResp.NickName)

	// 消息循环
	//go func() {
	for {
		// 读取消息
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			// 客户端断开连接或读取错误
			fmt.Printf("Read error:", err)
			break
		}
		handleMessage(message, messageType, uid)
	}
	//}()

	resp = &types.OnLineResponse{
		Status: true,
	}
	return
}

func handleMessage(message []byte, messageType int, senderUid int) {
	fmt.Printf("Received: %s (Type: %d)", message, messageType)

	chatReq := ChatRequest{}
	json.Unmarshal(message, &chatReq)
	fmt.Printf("Received to_uid: %s, msg: %s", chatReq.RevUserID, chatReq.Msg.Content)
	// 消息入库持久化

	//如果目标用户在线,转发该消息
	sendUserWs, ok := UserWsMap[uint(senderUid)]
	recvUid, _ := strconv.Atoi(chatReq.RevUserID)
	revUserWs, ok := UserWsMap[uint(recvUid)]
	if ok {
		resp := ChatResponse{
			RevUser:  revUserWs.UserInfo,
			SendUser: sendUserWs.UserInfo,
			Msg:      chatReq.Msg,
			CreateAt: time.Now(),
		}
		byteData, _ := json.Marshal(resp)
		// 写入消息 (回复给客户端)
		if err := revUserWs.Conn.WriteMessage(websocket.TextMessage, byteData); err != nil {
			fmt.Println("Write error:", err)
			panic(err)
		}
	}
}

type ChatRequest struct {
	RevUserID string    `json:"rev_user_id"`
	Msg       ctype.Msg `json:"msg"`
}

type ChatResponse struct {
	RevUser  UserInfo  `json:"rev_user"`
	SendUser UserInfo  `json:"send_user"`
	Msg      ctype.Msg `json:"msg"`
	CreateAt time.Time `json:"create_at"`
}
