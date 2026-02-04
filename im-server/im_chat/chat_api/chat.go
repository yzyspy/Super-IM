// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/netx"
	"github.com/zeromicro/go-zero/rest"
	"im-server/core"
	"im-server/im_chat/chat_api/internal/config"
	"im-server/im_chat/chat_api/internal/handler"
	"im-server/im_chat/chat_api/internal/svc"
)

var configFile = flag.String("f", "etc/chat.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// http服务地址注册到etcd，网关服务服务发现，从etcd获取具体到服务地址和端口
	httpApiUrl := fmt.Sprintf("http://%s:%d", netx.InternalIp(), c.Port)
	core.PutKv(ctx.Etcd, "chat_api", httpApiUrl)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
	fmt.Printf("Starting server success %s:%d...\n", c.Host, c.Port)
}

//
//// CORS 中间件，允许客户端连接
//func corsMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// 允许所有来源，以方便本地 HTML 文件访问
//		w.Header().Set("Access-Control-Allow-Origin", "*")
//		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
//		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
//
//		if r.Method == "OPTIONS" {
//			w.WriteHeader(http.StatusOK)
//			return
//		}
//		next.ServeHTTP(w, r)
//	})
//}
//
//// 处理新的 WebSocket 连接
//func wsHandler(ws *websocket.Conn) {
//	mutex.Lock()
//
//	// 尝试将连接分配给 A 或 B
//	var currentClient **websocket.Conn // 指向 clientA 或 clientB 的指针
//	var clientName string
//
//	if clientA == nil {
//		clientA = ws
//		currentClient = &clientA
//		clientName = "User A"
//		log.Println("New connection established as User A.")
//	} else if clientB == nil {
//		clientB = ws
//		currentClient = &clientB
//		clientName = "User B"
//		log.Println("New connection established as User B. Chat room is full.")
//		// 通知 User A 对方已连接
//		sendMessageToOther(clientB, "系统消息：User B 已加入聊天。")
//		// 通知 User B 对方已连接
//		sendMessageToOther(clientA, "系统消息：User A 已加入聊天。")
//	} else {
//		// 房间满员，拒绝连接
//		websocket.Message.Send(ws, "系统消息：聊天室已满，请稍后再试。")
//		ws.Close()
//		mutex.Unlock()
//		return
//	}
//
//	mutex.Unlock()
//
//	// 循环读取消息
//	for {
//		var message string
//		// 阻塞等待接收消息
//		if err := websocket.Message.Receive(ws, &message); err != nil {
//			log.Printf("%s connection closed or error: %v", clientName, err)
//			break
//		}
//
//		log.Printf("[%s] Received: %s", clientName, message)
//
//		// 转发消息给另一个客户端
//		sendMessageToOther(ws, fmt.Sprintf("%s: %s", clientName, message))
//	}
//
//	// --- 连接关闭后清理 ---
//	mutex.Lock()
//	if *currentClient == ws {
//		*currentClient = nil
//		log.Printf("%s disconnected. Slot cleared.", clientName)
//		// 通知另一个用户连接断开
//		sendMessageToOther(ws, fmt.Sprintf("系统消息：%s 已断开连接。", clientName))
//	}
//	mutex.Unlock()
//}
//
//// 存储两个客户端连接
//var (
//	clientA *websocket.Conn
//	clientB *websocket.Conn
//	mutex   sync.Mutex
//)
//
//// 发送消息给另一个客户端
//func sendMessageToOther(sender *websocket.Conn, message string) {
//	// mutex.Lock()
//	// defer mutex.Unlock()
//
//	var receiver *websocket.Conn
//
//	if sender == clientA && clientB != nil {
//		receiver = clientB
//	} else if sender == clientB && clientA != nil {
//		receiver = clientA
//	}
//
//	if receiver != nil {
//		if err := websocket.Message.Send(receiver, message); err != nil {
//			log.Printf("Error sending message to receiver: %v", err)
//		}
//		fmt.Printf("转发消息成功")
//	} else {
//		log.Println("Receiver not available or room is empty.")
//	}
//}
