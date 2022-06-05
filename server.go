package zdpgo_cache_http

import "github.com/zhangdapeng520/zdpgo_log"

/*
@Time : 2022/6/5 14:13
@Author : 张大鹏
@File : server
@Software: Goland2021.3.1
@Description:
*/

type Server struct {
	ICache
	Log *zdpgo_log.Log
}

// NewServer 后期可以基于此方法定制缓存方式
func NewServer(c ICache, log *zdpgo_log.Log) *Server {
	return &Server{
		ICache: c,
		Log:    log,
	}
}
