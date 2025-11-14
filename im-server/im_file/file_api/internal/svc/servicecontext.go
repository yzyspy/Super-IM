// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"im-server/core"
	"im-server/im_file/file_api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Etcd   *clientv3.KV
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Etcd:   core.InitEtcd(c.Etcd.Host, c.Etcd.Port),
	}
}
