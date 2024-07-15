package routers

import (
	"github.com/gin-gonic/gin"
	"project_05/middleware"
	"project_05/web"
)

// InitCourseRouter
// @Description 初始化课程路由
// @Author Oberl-Fitzgerald 2024-07-15 17:32:44
// @Param  r *gin.Engine
func InitCourseRouter(r *gin.Engine) {
	//初始化课程路由，使用中间件
	v1 := r.Group("/v1", middleware.TokenCheck, middleware.AuthCheck)
	v1.GET("/course", web.GetCourse)
	v1.POST("/course", web.CreateCourse)
	v1.PUT("/course", web.UpdateCourse)
	v1.DELETE("/course", web.DeleteCourse)

}
