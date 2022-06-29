package zdpgo_cache_http

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_requests"
)

/*
@Time : 2022/6/5 14:47
@Author : 张大鹏
@File : client
@Software: Goland2021.3.1
@Description:
*/

type Client struct {
	Config   *Config
	Requests *zdpgo_requests.Requests
	BaseUrl  string
}

func NewClient(req *zdpgo_requests.Requests, config *Config) *Client {
	addr := fmt.Sprintf("http://%s:%d", config.Client.Host, config.Client.Port)
	return &Client{
		Requests: req,
		Config:   config,
		BaseUrl:  addr,
	}
}

// GetStat 获取缓存状态信息
func (c *Client) GetStat() *Stat {
	response, err := c.Requests.Get(c.BaseUrl + "/status")
	if err != nil {
		return nil
	}

	// 解析json数据
	var stat Stat
	err = json.Unmarshal(response.Content, &stat)
	if err != nil {
		return nil
	}

	// 返回
	return &stat
}

// Set 设置键值对
func (c *Client) Set(key, value string) bool {
	response, err := c.Requests.Post(c.BaseUrl+"/cache/"+key, value)
	if err != nil {
		return false
	}

	// 判断是否添加成功
	flag := response.StatusCode == 200

	// 返回
	return flag
}

// Get 根据键获取值
func (c *Client) Get(key string) string {
	response, err := c.Requests.Get(c.BaseUrl + "/cache/" + key)
	if err != nil {
		return ""
	}

	// 返回
	return response.Text
}

// Delete 根据键删除值，不关心删除结果
func (c *Client) Delete(key string) {
	_, err := c.Requests.Delete(c.BaseUrl + "/cache/" + key)
	if err != nil {
		return
	}
}
