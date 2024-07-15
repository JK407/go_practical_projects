package web

import "github.com/gin-gonic/gin"

// Ping
// @Description ping 请求
// @Author Oberl-Fitzgerald 2024-07-15 17:33:09
// @Param  c *gin.Context
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"status":  200,
	})
}
