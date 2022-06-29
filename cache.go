package zdpgo_cache_http

/*
@Time : 2022/6/5 14:05
@Author : 张大鹏
@File : cache
@Software: Goland2021.3.1
@Description:
*/

import (
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
	"net/http"
)

type Cache struct {
	Config *Config
	Server *Server
	Client *Client
}

func New() *Cache {
	return NewWithConfig(&Config{})
}

func NewWithConfig(config *Config) *Cache {
	c := &Cache{}

	// 配置
	if config.Server.Host == "" {
		config.Server.Host = "0.0.0.0"
	}
	if config.Server.Port == 0 {
		config.Server.Port = 36333
	}
	if config.Client.Host == "" {
		config.Client.Host = "127.0.0.1"
	}
	if config.Client.Port == 0 {
		config.Client.Port = 36333
	}
	c.Config = config

	// 服务
	c.Server = NewServer(NewMemoryCache())

	// 客户端
	requests := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		UserAgent: "https://github.com/zhangdapeng520/zdpgo_cache_http",
	})
	c.Client = NewClient(requests, config)

	// 返回
	return c
}

func (c *Cache) Run() error {
	// 添加路由
	http.Handle("/cache/", c.Server.cacheHandler())
	http.Handle("/status", c.Server.statusHandler())

	// 启动服务
	address := fmt.Sprintf("%s:%d", c.Config.Server.Host, c.Config.Server.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		return err
	}
	return nil
}

// GetClient 获取客户端
func (c *Cache) GetClient() *Client {
	// 如果已存在，则直接返回
	if c.Client != nil {
		return c.Client
	}

	// 创建客户端
	requests := zdpgo_requests.NewWithConfig(&zdpgo_requests.Config{
		UserAgent: "https://github.com/zhangdapeng520/zdpgo_cache_http",
	})
	c.Client = NewClient(requests, c.Config)

	// 返回
	return c.Client
}
