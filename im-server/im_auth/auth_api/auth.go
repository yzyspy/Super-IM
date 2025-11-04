// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package main

import (
	"flag"
	"fmt"
	"im-server/core"

	"im-server/im_auth/auth_api/internal/config"
	"im-server/im_auth/auth_api/internal/handler"
	"im-server/im_auth/auth_api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/auth.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors()) // 配置可以跨域访问
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// http服务地址注册到etcd，网关服务服务发现，从etcd获取具体到服务地址和端口
	httpApiUrl := fmt.Sprintf("http://%s:%d", c.Host, c.Port)
	core.PutKv(ctx.Etcd, "auth_api", httpApiUrl)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
