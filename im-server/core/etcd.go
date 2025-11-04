package core

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

//go get go.etcd.io/etcd/client/v3

// 127.0.0.1:2379
func InitEtcd(host string, port int) *clientv3.KV {
	// expect dial time-out on ipv4 blackhole
	serverUrl := fmt.Sprintf("http://%s:%d", host, port)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{serverUrl}, // etcd 服务器地址和端口
		DialTimeout: 2 * time.Second,
	})

	// etcd clientv3 >= v3.2.10, grpc/grpc-go >= v1.7.3
	if err == context.DeadlineExceeded {
		// handle errors
	}

	kv := clientv3.NewKV(cli)

	return &kv
}

func PutKv(kv *clientv3.KV, key string, value string) {
	(*kv).Put(context.TODO(), key, value)
	fmt.Println("Put key-value pair to etcd successfully." + key + " : " + value)
}

func GetKv(kv *clientv3.KV, key string) string {
	resp, error := (*kv).Get(context.TODO(), key)
	if error != nil {
		fmt.Println("Get key-value pair from etcd failed." + key)
		return ""
	}
	return string(resp.Kvs[0].Value)
}
