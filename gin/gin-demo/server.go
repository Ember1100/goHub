package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/pragmaticreviews/golang-gin-poc/controller"
	"gitlab.com/pragmaticreviews/golang-gin-poc/service"
)

var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "OK",
		})
	})
	server.POST("videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.Save(ctx))
	})
	server.GET("videos", func(ctx *gin.Context) {
		ctx.JSON(200, VideoController.FindAll())
	})

	server.Run(":8080")
}
