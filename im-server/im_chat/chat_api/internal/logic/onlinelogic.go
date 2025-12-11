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
	upgrade, err := upgrader.Upgrade(w, r, nil)
	upgrade.ReadMessage()

	resp = &types.OnLineResponse{
		Status: true,
	}
	return
}
