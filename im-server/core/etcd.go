package core

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

//go get go.etcd.io/etcd/client/v3

// 127.0.0.1:2379
func InitEtcd() clientv3.KV {
	// expect dial time-out on ipv4 blackhole
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://127.0.0.1:2379"}, // etcd 服务器地址和端口
		DialTimeout: 2 * time.Second,
	})

	// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
	if err == context.DeadlineExceeded {
		// handle errors
	}

	kv := clientv3.NewKV(cli)

	return kv
}
