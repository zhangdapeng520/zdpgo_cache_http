package zdpgo_cache_http

/*
@Time : 2022/6/5 14:26
@Author : 张大鹏
@File : icache
@Software: Goland2021.3.1
@Description: 缓存服务接口
*/

// ICache 缓存接口
type ICache interface {
	Set(string, []byte) error   // 设置值
	Get(string) ([]byte, error) // 获取值
	Del(string) error           // 删除键值对
	GetStat() Stat              // 获取缓存状态
}
