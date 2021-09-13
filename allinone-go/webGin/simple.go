package webGin
import (
	 "github.com/gin-gonic/gin"
)


func SimpleServer(){
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	r.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务

}
