package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	//gindump "github.com/tpkeeper/gin-dump"
	"github.com/valekar/gin-gonic/controller"
	"github.com/valekar/gin-gonic/middlewares"
	"github.com/valekar/gin-gonic/service"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	server.Use(middlewares.BasicAuth())
	//server.Use(gindump.Dump())

	server.GET("/videos/", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos/", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(200, gin.H{"ok": "ok"})
		}
	})

	server.Run(":8080")
}
