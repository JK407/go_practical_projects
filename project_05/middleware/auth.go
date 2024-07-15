package middleware

import "github.com/gin-gonic/gin"

// AuthCheck
// @Description 认证检查
// @Author Oberl-Fitzgerald 2024-07-15 17:30:50
// @Param  c *gin.Context
func AuthCheck(c *gin.Context) {
	//  获取用户信息
	userName, _ := c.Get("user_name")
	userId, _ := c.Get("user_id")
	//  用户信息为空，返回401
	if userName == "" || userId == 0 {
		c.JSON(401, gin.H{
			"message": "auth error",
			"status":  401,
		})
		//  终止请求
		c.Abort()
		return

	}
	//  继续请求
	c.Next()
}

// token
var token = "123456"

// TokenCheck
// @Description token检查
// @Author Oberl-Fitzgerald 2024-07-15 17:31:30
// @Param  c *gin.Context
func TokenCheck(c *gin.Context) {
	//获取请求头中的token
	accessToken := c.Request.Header.Get("access_token")
	//  token不一致，返回401
	if accessToken != token {
		c.JSON(401, gin.H{
			"message": "token error",
			"status":  401,
		})
		c.Abort()
		return
	}
	//  设置用户信息
	c.Set("user_name", "oberl")
	c.Set("user_id", 1001)
	//  继续请求
	c.Next()
}
