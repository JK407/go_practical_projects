package web

import "github.com/gin-gonic/gin"

// LoginReq
// @Description: 登录请求
// @Author Oberl-Fitzgerald 2024-07-15 17:32:54
type LoginReq struct {
	UserName string `form:"user_name"`
	Password string `form:"password"`
}

// Login
// @Description 登录
// @Author Oberl-Fitzgerald 2024-07-15 17:32:57
// @Param  c *gin.Context
func Login(c *gin.Context) {
	req := &LoginReq{}
	//  绑定请求
	c.Bind(req)
	c.JSON(200, gin.H{
		"message": "login",
		"status":  200,
		"data":    req,
	})
}
