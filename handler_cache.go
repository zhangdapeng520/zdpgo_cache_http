package zdpgo_cache_http

import (
	"github.com/zhangdapeng520/zdpgo_log"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
@Time : 2022/6/5 14:29
@Author : 张大鹏
@File : handler_cache
@Software: Goland2021.3.1
@Description: 缓存处理器
*/

type cacheHandler struct {
	*Server
	Log *zdpgo_log.Log
}

// ServeHTTP 实现ServeHTTP接口
func (h *cacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 获取key
	key := strings.Split(r.URL.EscapedPath(), "/")[2]
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 获取方法
	m := r.Method

	// 新增或修改值
	if m == http.MethodPost {
		b, _ := ioutil.ReadAll(r.Body)
		if len(b) != 0 {
			e := h.Set(key, b)
			if e != nil {
				h.Log.Error("新增或修改缓存数失败", "error", e)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
		return
	}

	// 获取值
	if m == http.MethodGet {
		b, e := h.Get(key)
		if e != nil {
			h.Log.Error("根据键获取缓存的值失败", "error", e)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if len(b) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Write(b)
		return
	}

	// 删除值
	if m == http.MethodDelete {
		e := h.Del(key)
		if e != nil {
			h.Log.Error("根据键删除值失败", "error", e)
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	// 返回响应
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// 缓存处理器
func (s *Server) cacheHandler() http.Handler {
	return &cacheHandler{
		Server: s,
		Log:    s.Log,
	}
}
