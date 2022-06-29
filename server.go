package zdpgo_cache_http

/*
@Time : 2022/6/5 14:13
@Author : 张大鹏
@File : server
@Software: Goland2021.3.1
@Description:
*/

type Server struct {
	ICache
}

// NewServer 后期可以基于此方法定制缓存方式
func NewServer(c ICache) *Server {
	return &Server{
		ICache: c,
	}
}
