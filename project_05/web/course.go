package web

import "github.com/gin-gonic/gin"

// CreateCourse
// @Description 创建课程
// @Author Oberl-Fitzgerald 2024-07-15 17:32:18
// @Param  c *gin.Context
func CreateCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "create course",
		"status":  200,
	})
}

// UpdateCourse
// @Description 更新课程
// @Author Oberl-Fitzgerald 2024-07-15 17:32:23
// @Param  c *gin.Context
func UpdateCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "update course",
		"status":  200,
	})
}

// DeleteCourse
// @Description 删除课程
// @Author Oberl-Fitzgerald 2024-07-15 17:32:26
// @Param  c *gin.Context
func DeleteCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "delete course",
		"status":  200,
	})
}

// GetCourse
// @Description 获取课程
// @Author Oberl-Fitzgerald 2024-07-15 17:32:34
// @Param  c *gin.Context
func GetCourse(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "get course",
		"status":  200,
	})
}
