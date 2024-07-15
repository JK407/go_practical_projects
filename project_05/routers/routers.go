package routers

import "github.com/gin-gonic/gin"

// InitRouter
// @Description 初始化路由
// @Author Oberl-Fitzgerald 2024-07-15 17:32:40
// @Param  r *gin.Engine
func InitRouter(r *gin.Engine) {
	//初始化课程路由
	InitCourseRouter(r)
	//初始化 api 路由
	InitApiRouter(r)
	//  运行端口
	r.Run(":8080")
}
