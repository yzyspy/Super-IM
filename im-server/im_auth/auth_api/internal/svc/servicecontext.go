// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/go-redis/redis"
	"im-server/core"
	"im-server/im_auth/auth_api/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
	Redis  *redis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     core.InitGorm(c.MySql.DataSource),
		Redis:  core.InitRedis(c.Redis.Host, c.Redis.Port),
	}
}
