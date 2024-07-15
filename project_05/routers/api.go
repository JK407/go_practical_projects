package routers

import (
	"github.com/gin-gonic/gin"
	"project_05/web"
)

// InitApiRouter
// @Description 初始化 api 路由
// @Author Oberl-Fitzgerald 2024-07-15 17:32:47
// @Param  r *gin.Engine
func InitApiRouter(r *gin.Engine) {
	//初始化 api 路由
	api := r.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", web.Ping)
	v1.POST("/login", web.Login)
	v1.POST("/register", web.Register)
}
