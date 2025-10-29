// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/zrpc"
	"im-server/core"
	"im-server/im_auth/auth_api/internal/config"
	"im-server/im_user/user_rpc/types/user_rpc"
	"im-server/im_user/user_rpc/user"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gorm.DB
	Redis   *redis.Client
	UserRpc user_rpc.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DB:      core.InitGorm(c.MySql.DataSource),
		Redis:   core.InitRedis(c.Redis.Host, c.Redis.Port),
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)), //创建rpc客户端
	}
}
