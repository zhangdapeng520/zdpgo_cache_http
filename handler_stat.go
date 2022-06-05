package zdpgo_cache_http

import (
	"encoding/json"
	"log"
	"net/http"
)

/*
@Time : 2022/6/5 14:28
@Author : 张大鹏
@File : handler_stat
@Software: Goland2021.3.1
@Description: 状态处理器
*/

// 状态处理器
type statusHandler struct {
	*Server
}

// ServeHTTP 实现监听HTTP的方法
func (h *statusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	b, e := json.Marshal(h.GetStat())
	if e != nil {
		log.Println(e)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(b)
}

// 状态处理器
func (s *Server) statusHandler() http.Handler {
	return &statusHandler{s}
}
