package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"project_06/breaker"
	"project_06/limiter"
	"project_06/middleware"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/ping", middleware.Limiter(limiter.NewLimiter(2*time.Second, 5)), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})
	})
	//  初始化熔断器，连续失败次数4次，连续成功次数4次，半开状态下最大请求数2次，熔断器时间周期15秒
	b := breaker.NewBreaker(4, 4, 2, time.Second*15)
	r.POST("/ping", func(c *gin.Context) {
		err := b.Exec(func() error {
			val, _ := c.GetQuery("value")
			if val == "a" {
				return errors.New("value is a")
			}
			return nil
		})
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status":  500,
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "pong",
			"status":  200,
		})

	})
	r.Run()
}
