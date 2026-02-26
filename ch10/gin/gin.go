package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// go modules  添加依赖、依赖分析、删除未使用的依赖项
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()

}
