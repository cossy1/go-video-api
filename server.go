package main

import (
	"go-api/controller"
	"go-api/middlewares"
	"go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.NewVideoService()
	videoController controller.VideoController = controller.NewVideoController(videoService)
)

func main() {
	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, videoController.FindAll())

	})

	server.POST("/videos", func(ctx *gin.Context) {
		videoController.SaveVideo(ctx)
	})

	server.Run(":8080")
}
