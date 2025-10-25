// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package config

import "github.com/zeromicro/go-zero/rest"

//type MySql struct {
//	DataSource string
//}
//type Config struct {
//	rest.RestConf
//	MySql
//}

type Config struct {
	rest.RestConf
	MySql struct {
		DataSource string
	}

	Redis struct {
		Host     string
		Port     int
		Password string
	}
}
