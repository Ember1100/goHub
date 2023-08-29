package main

import (
	"go-gin-example/conf"
	"go-gin-example/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 禁用控制台颜色，将日志写入文件时不需要控制台颜色。
	// gin.DisableConsoleColor()

	// 记录到文件。
	// f, _ := os.Create("gin.log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// 如果需要同时将日志写入文件和控制台，请使用以下代码。
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	router := routers.SetupRouter()

	// 在中间件中设置全局数据库连接
	router.Use(func(c *gin.Context) {
		db, err := conf.ConnectToDatabase()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		c.Set("db", db)
		c.Next()
	})

	router.Run(":3080")

}
