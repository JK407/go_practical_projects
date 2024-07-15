package main

import (
	"github.com/gin-gonic/gin"
	"project_05/routers"
)

// main
// @Description 主函数
// @Author Oberl-Fitzgerald 2024-07-15 17:34:06
func main() {
	//  初始化路由
	r := gin.Default()
	routers.InitRouter(r)
}
