package main

import (
	"fmt"
	"project_01/server"
)

func main() {
	cache := server.NewCacheService()
	cache.SetMaxMemory("100MB")
	cache.Set("name", "oberl")
	cache.Set("age", 18)
	fmt.Println(cache.GetCache())
}
