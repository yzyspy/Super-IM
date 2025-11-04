package main

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"im-server/core"
	"io"
	"log"
	"net/http"
	"strings"
)

var serviceMap = map[string]string{
	"user": "http://localhost:9001", //user服务http服务端地址和端口
	"auth": "http://localhost:9002", //auth服务http服务端地址和端口
}

// 定义处理函数
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	//请求认证服务,认证通过直接转发请求到下游服务，认证不通过，直接返回需要登录

	authServerAddr := core.GetKv(etcd, "auth_api")

	var newRequest *http.Request
	if strings.Contains(r.URL.Path, "api/auth") {
		// 处理api/auth请求
		newRequest, _ = http.NewRequest(r.Method, serviceMap["auth"]+r.URL.Path, r.Body)
	} else if strings.Contains(r.URL.Path, "api/user") {
		// 处理api/user
	}
	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//这一步很重要
	io.Copy(w, resp.Body)

}

var etcd *clientv3.KV

func main() {
	// 注册处理函数到特定路由
	http.HandleFunc("/", redirectHandler)

	etcd = core.InitEtcd("127.0.0.1", 2379)

	// 启动服务器，监听8080端口
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
