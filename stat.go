package zdpgo_cache_http

/*
@Time : 2022/6/5 14:28
@Author : 张大鹏
@File : stat
@Software: Goland2021.3.1
@Description: 缓存状态
*/

// Stat 状态
type Stat struct {
	Count     int64 `json:"count"`      // 存放的数量
	KeySize   int64 `json:"key_size"`   // key大小
	ValueSize int64 `json:"value_size"` // 值大小
}

// add 添加状态
func (s *Stat) add(k string, v []byte) {
	s.Count += 1
	s.KeySize += int64(len(k))
	s.ValueSize += int64(len(v))
}

// del 删除状态
func (s *Stat) del(k string, v []byte) {
	s.Count -= 1
	s.KeySize -= int64(len(k))
	s.ValueSize -= int64(len(v))
}
