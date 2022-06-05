package main

import (
	"github.com/zhangdapeng520/zdpgo_cache_http"
)

/*
@Time : 2022/6/5 14:04
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description: HTTP缓存服务
*/

func main() {
	cache := zdpgo_cache_http.NewWithConfig(&zdpgo_cache_http.Config{
		Debug: true,
		Server: zdpgo_cache_http.HttpInfo{
			Port: 3333,
		},
	})
	err := cache.Run()
	if err != nil {
		panic(err)
	}
}
