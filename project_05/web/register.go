package web

import "github.com/gin-gonic/gin"

// RegisterReq
// @Description: 注册请求
// @Author Oberl-Fitzgerald 2024-07-15 17:33:18
type RegisterReq struct {
	UserName string `form:"user_name" binding:"required"`    // UserName @Description: 用户名，必填
	Password string `form:"password" binding:"required"`     // Password @Description: 密码，必填
	Phone    string `form:"phone" binding:"required,e164"`   // Phone @Description: 手机号，必填，e164 格式（校验手机号）
	Email    string `form:"email" binding:"omitempty,email"` // Email @Description: 邮箱，非必填，邮箱格式（校验邮箱）
}

// Register
// @Description 注册
// @Author Oberl-Fitzgerald 2024-07-15 17:33:51
// @Param  c *gin.Context
func Register(c *gin.Context) {
	req := &RegisterReq{}
	//  绑定请求
	err := c.ShouldBind(req)
	//判断是否绑定成功
	if err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
			"status":  400,
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "register",
		"status":  200,
		"data":    req,
	})
}
