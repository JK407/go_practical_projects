package cache

import (
	"encoding/json"
	"fmt"
	"log"
	"project_01/utils"
	"sync"
	"time"
)

// memCache
// @Description: 内存缓存
// @Author Oberl-Fitzgerald 2024-07-08 15:48:40
type memCache struct {
	maxCacheSize        int64                   // maxCacheSize @Description: 缓存最大内存
	maxCacheSizeStr     string                  // maxCacheSizeStr @Description: 缓存最大内存字符串
	currentCacheSize    int64                   // currentCacheSize @Description: 当前缓存内存
	currentCacheSizeStr string                  // currentCacheSizeStr @Description: 当前缓存内存字符串
	memCacheValues      map[string]*memCacheVal // memCacheValues @Description: 缓存值
	locker              sync.RWMutex            // locker @Description: 读写锁
	clearExpiredCache   time.Duration           // clearExpiredCache @Description: 清除过期缓存
}

// memCacheVal
// @Description: 缓存值
// @Author Oberl-Fitzgerald 2024-07-08 15:48:44
type memCacheVal struct {
	Value  interface{} `json:"value"`  // value @Description: 缓存值
	Expire time.Time   `json:"expire"` // expire @Description: 过期时间
	Size   int64       `json:"size"`   // size @Description: 缓存大小
}

// NewMemCache
// @Description 创建一个新的内存缓存
// @Author Oberl-Fitzgerald 2024-07-08 15:48:48
// @Return Cache
func NewMemCache() Cache {
	mc := &memCache{
		//  初始化缓存值
		memCacheValues:    make(map[string]*memCacheVal, 0),
		clearExpiredCache: time.Second * 3,
	}
	go mc.ClearExpiredCache()
	return mc
}

// SetMaxMemory
// @Description 设置缓存最大内存
// @Author Oberl-Fitzgerald 2024-07-08 15:01:19
// @Param  size string
// @Return bool
func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxCacheSize, mc.maxCacheSizeStr = utils.ParseSize(size)
	//logStr := fmt.Sprintf("call [method:%s] success, [result:%#v]", "SetMaxMemory", mc)
	//log.Println(logStr)
	return true
}

// Set
// @Description 将value写入缓存
// @Author Oberl-Fitzgerald 2024-07-08 15:08:33
// @Param  key string
// @Param  value interface{}
// @Param  expire time.Duration
// @Return bool
func (mc *memCache) Set(key string, value interface{}, expire time.Duration) bool {
	//  加写锁
	mc.locker.Lock()
	defer mc.locker.Unlock()
	v := &memCacheVal{
		Value:  value,
		Expire: time.Now().Add(expire),
		Size:   utils.GetValueSize(value),
	}
	//  如果缓存值存在，则删除
	mc.del(key)
	//  添加缓存值
	mc.add(key, v)
	//  如果当前缓存内存大于最大缓存内存，则删除缓存值
	if mc.currentCacheSize > mc.maxCacheSize {
		mc.del(key)
		log.Fatal(fmt.Sprintf("call [method:%s] failed, [result: CurrentCacheSize > MaxCacheSize ]", "Set"))
	}
	//logStr := fmt.Sprintf("call [method:%s] success, [result:%#v]", "Set", mc)
	//log.Println(logStr)
	return true
}

// Get
// @Description 根据key值获取value
// @Author Oberl-Fitzgerald 2024-07-08 15:47:56
// @Param  key string
// @Return interface{}
// @Return bool
func (mc *memCache) Get(key string) (interface{}, bool) {
	//  加读锁
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	//  如果缓存值存在，则判断是否过期
	if v, ok := mc.get(key); ok {
		if time.Now().After(v.Expire) {
			//  如果过期，则删除缓存值
			mc.del(key)
			return nil, false
		}
		return v.Value, ok
	}
	return nil, false
}

// get
// @Description 获取缓存值
// @Author Oberl-Fitzgerald 2024-07-08 15:50:36
// @Param  key string
// @Return *memCacheVal
// @Return bool
func (mc *memCache) get(key string) (*memCacheVal, bool) {
	v, ok := mc.memCacheValues[key]
	return v, ok
}

// del
// @Description 删除缓存值
// @Author Oberl-Fitzgerald 2024-07-08 15:50:43
// @Param  key string
func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		//  释放内存
		mc.currentCacheSize -= tmp.Size
		//  删除缓存值
		delete(mc.memCacheValues, key)

	}
}

// add
// @Description 添加缓存值
// @Author Oberl-Fitzgerald 2024-07-08 15:50:49
// @Param  key string
// @Param  val *memCacheVal
func (mc *memCache) add(key string, val *memCacheVal) {
	mc.memCacheValues[key] = val
	mc.currentCacheSize += val.Size
}

// Del
// @Description 删除key值
// @Author Oberl-Fitzgerald 2024-07-08 15:52:07
// @Param  key string
// @Return bool
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	fmt.Println("call Del")
	mc.del(key)
	return true
}

// Exists
// @Description 判断key是否存在
// @Author Oberl-Fitzgerald 2024-07-08 15:52:19
// @Param  key string
// @Return bool
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	fmt.Println("call Exists")
	_, ok := mc.get(key)
	return ok
}

// Flush
// @Description 清空所有key
// @Author Oberl-Fitzgerald 2024-07-08 15:52:27
// @Return bool
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	fmt.Println("call Flush")
	//  清空缓存值
	mc.memCacheValues = make(map[string]*memCacheVal, 0)
	mc.currentCacheSize = 0
	return true
}

// Keys
// @Description 获取缓存中所有key的数量
// @Author Oberl-Fitzgerald 2024-07-08 15:53:12
// @Return int64
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	fmt.Println("call Keys")
	return int64(len(mc.memCacheValues))
}

// ClearExpiredCache
// @Description 清除过期缓存
// @Author Oberl-Fitzgerald 2024-07-08 16:03:24
func (mc *memCache) ClearExpiredCache() {
	timeTicker := time.NewTicker(mc.clearExpiredCache)
	defer timeTicker.Stop()
	for {
		select {
		//  定时清除过期缓存
		case <-timeTicker.C:
			mc.locker.Lock()
			for key, val := range mc.memCacheValues {
				if time.Now().After(val.Expire) {
					mc.del(key)
				}
			}
			mc.locker.Unlock()
		}
	}
}

// GetCache
// @Description 获取全部缓存，返回json字符串
// @Author Oberl-Fitzgerald 2024-07-08 16:38:58
// @Return string
func (mc *memCache) GetCache() string {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	cacheValues := make(map[string]interface{})
	for key, val := range mc.memCacheValues {
		cacheValues[key] = val
	}
	jsonData, err := json.Marshal(cacheValues)
	if err != nil {
		log.Fatal("Error marshaling cache values:", err)
		return ""
	}
	return string(jsonData)
}
