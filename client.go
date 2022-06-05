package zdpgo_cache_http

import (
	"encoding/json"
	"fmt"
	"github.com/zhangdapeng520/zdpgo_log"
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
	Log      *zdpgo_log.Log
	BaseUrl  string
}

func NewClient(req *zdpgo_requests.Requests, config *Config) *Client {
	addr := fmt.Sprintf("http://%s:%d", config.Client.Host, config.Client.Port)
	return &Client{
		Requests: req,
		Log:      req.Log,
		Config:   config,
		BaseUrl:  addr,
	}
}

// GetStat 获取缓存状态信息
func (c *Client) GetStat() *Stat {
	response, err := c.Requests.Any(zdpgo_requests.Request{
		Method: "GET",
		Url:    c.BaseUrl + "/status",
	})
	if err != nil {
		c.Log.Error("请求缓存状态失败", "error", err)
		return nil
	}

	// 解析json数据
	var stat Stat
	err = json.Unmarshal(response.Content, &stat)
	if err != nil {
		c.Log.Error("解析状态信息失败", "error", err)
		return nil
	}

	// 返回
	return &stat
}

// Set 设置键值对
func (c *Client) Set(key, value string) bool {
	response, err := c.Requests.AnyText(zdpgo_requests.Request{
		Method: "POST",
		Url:    c.BaseUrl + "/cache/" + key,
		Text:   value,
	})
	if err != nil {
		c.Log.Error("设置缓存键值对失败", "error", err)
		return false
	}

	// 判断是否添加成功
	flag := response.StatusCode == 200

	// 返回
	return flag
}

// Get 根据键获取值
func (c *Client) Get(key string) string {
	response, err := c.Requests.Any(zdpgo_requests.Request{
		Method: "GET",
		Url:    c.BaseUrl + "/cache/" + key,
	})
	if err != nil {
		c.Log.Error("根据键获取缓存键值失败", "error", err)
		return ""
	}

	// 返回
	return response.Text
}

// Delete 根据键删除值，不关心删除结果
func (c *Client) Delete(key string) {
	_, err := c.Requests.Any(zdpgo_requests.Request{
		Method: "DELETE",
		Url:    c.BaseUrl + "/cache/" + key,
	})
	if err != nil {
		c.Log.Error("根据键删除值失败", "error", err)
		return
	}
}
