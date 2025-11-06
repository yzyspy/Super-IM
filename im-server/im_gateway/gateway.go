package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"im-server/core"
	"io"
	"log"
	"net/http"
	"strings"
)

//var serviceMap = map[string]string{
//	"user": "http://localhost:9001", //user服务http服务端地址和端口
//	"auth": "http://localhost:9002", //auth服务http服务端地址和端口
//}

type AuthResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// 定义处理函数
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)

	//请求认证服务,认证通过直接转发请求到下游服务，认证不通过，直接返回需要登录
	authServerAddr := core.GetKv(etcd, "auth_api")

	authServerUrl := fmt.Sprintf("%s/api/auth/authentication", authServerAddr)

	log.Printf("redirectHandler authServerUrl=%s", authServerUrl)

	//请求认证服务
	authRequest, _ := http.NewRequest("POST", authServerUrl, bytes.NewBuffer(reqBody))
	authRequest.Header = r.Header
	authRequest.Header.Set("ValidPath", r.URL.Path)
	authRequest.Header.Set("X-Forwarded-For", r.RemoteAddr)
	authResponse, authError := http.DefaultClient.Do(authRequest)

	authResponseBytes, _ := io.ReadAll(authResponse.Body)
	var authResponseObj AuthResponse
	json.Unmarshal(authResponseBytes, &authResponseObj)

	fmt.Printf("redirectHandler 网关请求认证服务成功 request=%+v  authResponseObj=%+v", r, authResponseObj)
	if authResponseObj.Code != 0 {
		log.Printf("redirectHandler authError=%+v", authError)
		http.Error(w, errors.New("token异常请重新登录").Error(), http.StatusUnauthorized)
		return
	}

	// 认证通过，转发请求到下游服务
	var newRequest *http.Request
	if strings.Contains(r.URL.Path, "api/auth") {
		// 处理api/auth请求
		url := authServerAddr + r.URL.Path
		log.Printf("redirectHandler url=%s", url)
		newRequest, _ = http.NewRequest(r.Method, url, bytes.NewBuffer(reqBody))
		newRequest.Header = r.Header
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
