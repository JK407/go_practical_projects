package server

import (
	"project_01/cache"
	"time"
)

const (
	//  默认过期时间间隔
	DefaultExpire = time.Hour * 24
)

type CacheServer struct {
	memCache cache.Cache
}

func NewCacheService() *CacheServer {
	return &CacheServer{
		memCache: cache.NewMemCache(),
	}
}

func (cs *CacheServer) SetMaxMemory(size string) bool {
	return cs.memCache.SetMaxMemory(size)
}

func (cs *CacheServer) Set(key string, value interface{}, expire ...time.Duration) bool {
	var expireT = DefaultExpire
	if len(expire) > 0 {
		expireT = expire[0]
	}
	return cs.memCache.Set(key, value, expireT)
}

func (cs *CacheServer) Get(key string) (interface{}, bool) {
	return cs.memCache.Get(key)
}

func (cs *CacheServer) Del(key string) bool {
	return cs.memCache.Del(key)
}

func (cs *CacheServer) Exists(key string) bool {
	return cs.memCache.Exists(key)
}

func (cs *CacheServer) Flush() bool {
	return cs.memCache.Flush()
}

func (cs *CacheServer) Keys() int64 {
	return cs.memCache.Keys()
}

func (cs *CacheServer) GetCache() string {
	return cs.memCache.GetCache()
}
