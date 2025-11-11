package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"im-server/core"
	"io"
	"log"
	"net/http"
	"strings"
)

type AuthResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	UserId int64  `json:"user_id"`
	Role   int32  `json:"role"`
}

// 定义处理函数
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := io.ReadAll(r.Body)
	//请求认证服务,认证通过直接转发请求到下游服务，认证不通过，直接返回需要登录
	authRequest := makeAuthRequest(r, reqBody)
	authResponse, authError := http.DefaultClient.Do(authRequest)

	authResponseBytes, _ := io.ReadAll(authResponse.Body)
	var authResponseObj AuthResponse
	err := json.Unmarshal(authResponseBytes, &authResponseObj)
	if err != nil {
		fmt.Printf("redirectHandler 网关请求认证服务失败 err=%+v", err)
		return
	}

	fmt.Printf("redirectHandler 网关请求认证服务成功 request=%+v  authResponseObj=%+v", r, authResponseObj)
	if authResponseObj.Code != 0 {
		log.Printf("redirectHandler authError=%+v", authError)
		io.Copy(w, bytes.NewBuffer(authResponseBytes))
		//http.Error(w, errors.New("token异常请重新登录").Error(), http.StatusUnauthorized)
		return
	}
	// 认证通过，转发请求到下游服务
	newRequest := makeProxyRequest(r, reqBody)
	// 把登录用户的uid写进header
	newRequest.Header.Set("uid", fmt.Sprintf("%d", authResponseObj.UserId))

	resp, err := http.DefaultClient.Do(newRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//这一步很重要
	io.Copy(w, resp.Body)
}

func makeAuthRequest(r *http.Request, reqBody []byte) *http.Request {
	authServerAddr := core.GetKv(etcd, "auth_api")
	authServerUrl := fmt.Sprintf("%s/api/auth/authentication", authServerAddr)
	//请求认证服务
	authRequest, _ := http.NewRequest("POST", authServerUrl, bytes.NewBuffer(reqBody))
	authRequest.Header = r.Header
	authRequest.Header.Set("ValidPath", r.URL.Path)
	authRequest.Header.Set("X-Forwarded-For", r.RemoteAddr)
	return authRequest
}

func makeProxyRequest(r *http.Request, reqBody []byte) *http.Request {
	var newRequest *http.Request
	log.Printf("redirectHandler r.URL.Path=%s\n", r.URL.Path)
	var newUrl string
	if strings.Contains(r.URL.Path, "api/auth") {
		// 处理api/auth请求
		authServerAddr := core.GetKv(etcd, "auth_api")
		newUrl = authServerAddr + r.URL.Path
	} else if strings.Contains(r.URL.Path, "api/user") {
		// 处理api/user
		userServerAddr := core.GetKv(etcd, "user_api")
		newUrl = userServerAddr + r.URL.Path
	}
	newRequest, _ = http.NewRequest(r.Method, newUrl, bytes.NewBuffer(reqBody))
	newRequest.Header = r.Header
	newRequest.Header.Set("X-Forwarded-For", r.RemoteAddr)
	log.Printf("redirectHandler newRequest=%+v\n", newRequest)
	return newRequest
}

var etcd *clientv3.KV

func main() {
	http.HandleFunc("/", redirectHandler)
	etcd = core.InitEtcd("127.0.0.1", 2379)
	// 启动服务器，监听8080端口
	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
