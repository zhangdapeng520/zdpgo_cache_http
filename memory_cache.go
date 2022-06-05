package zdpgo_cache_http

import "sync"

/*
@Time : 2022/6/5 14:26
@Author : 张大鹏
@File : memory_cache
@Software: Goland2021.3.1
@Description: 基于内存的缓存服务
*/

// MemoryCache 内存缓存
type MemoryCache struct {
	c     map[string][]byte // 存放的数据池
	mutex sync.RWMutex      // 同步锁
	Stat                    // 状态
}

// Set 设置键值对
func (c *MemoryCache) Set(k string, v []byte) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// 如果存在，先删除存在的状态
	tmp, exist := c.c[k]
	if exist {
		c.del(k, tmp)
	}

	// 添加新的值和删除
	c.c[k] = v
	c.add(k, v)
	return nil
}

// Get 根据键获取值
func (c *MemoryCache) Get(k string) ([]byte, error) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.c[k], nil
}

// Del 根据键删除值
func (c *MemoryCache) Del(k string) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	// 如果存在则删除状态和键值对
	v, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, v)
	}
	return nil
}

// GetStat 获取状态
func (c *MemoryCache) GetStat() Stat {
	return c.Stat
}

// NewMemoryCache 创建内存缓存对象
func NewMemoryCache() *MemoryCache {
	return &MemoryCache{make(map[string][]byte), sync.RWMutex{}, Stat{}}
}
