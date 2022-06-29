package main

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_cache_http"
)

/*
@Time : 2022/6/5 14:46
@Author : 张大鹏
@File : main
@Software: Goland2021.3.1
@Description:
*/

func main() {
	cache := zdpgo_cache_http.NewWithConfig(&zdpgo_cache_http.Config{
		Debug: true,
		Client: zdpgo_cache_http.HttpInfo{
			Port: 3333,
		},
	})

	// 获取缓存信息
	client := cache.GetClient()
	stat := client.GetStat()
	fmt.Println(stat)

	// 设置键值对
	flag := client.Set("username", "zhangdapeng520")
	if !flag {
		panic("添加键值对失败")
	}
	stat = client.GetStat()
	fmt.Println(stat)

	// 根据键获取值
	username := client.Get("username")
	fmt.Println("根据键获取的值：", username)

	// 根据键删除值
	client.Delete("username")
	stat = client.GetStat()
	fmt.Println(stat)
}
