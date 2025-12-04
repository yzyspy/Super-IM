// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	clientv3 "go.etcd.io/etcd/client/v3"
	"gorm.io/gorm"
	"im-server/core"
	"im-server/im_chat/chat_api/internal/config"
	"im-server/im_user/user_rpc/types/user_rpc"
	"im-server/im_user/user_rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	Etcd    *clientv3.KV
	UserRpc user_rpc.UserClient
	DB      *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      core.InitGorm(c.MySql.DataSource),
		Etcd:    core.InitEtcd(c.Etcd.Host, c.Etcd.Port),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)), //创建rpc客户端
	}
}
