package server

import (
	"testing"
	"time"
)

func TestCacheServer(t *testing.T) {
	cache := NewCacheService()
	cache.SetMaxMemory("100MB")
	cache.Set("name", "oberl")
	cache.Set("age", 18)
	cache.Set("hobby", []string{"coding", "reading"}, time.Second*5)
	v, _ := cache.Get("hobby")
	t.Logf("hobby: %#v", v)
	t.Log(cache.GetCache())
	time.Sleep(time.Second * 6)
	t.Log(cache.GetCache())
}
