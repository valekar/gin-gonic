package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	//gindump "github.com/tpkeeper/gin-dump"
	"github.com/valekar/gin-gonic/controller"
	"github.com/valekar/gin-gonic/middlewares"
	"github.com/valekar/gin-gonic/repository"
	"github.com/valekar/gin-gonic/service"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()
	videoService    service.VideoService       = service.New(videoRepository)
	videoController controller.VideoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	defer videoRepository.CloseDB()

	setupLogOutput()

	server := gin.New()

	server.Static("css", "./templates/css")
	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery())
	server.Use(middlewares.Logger())
	//server.Use(middlewares.BasicAuth())
	//server.Use(gindump.Dump())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"ok": "ok"})
			}
		})

		apiRoutes.PUT("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Update(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"ok": "ok"})
			}
		})

		apiRoutes.DELETE("/videos/:id", func(ctx *gin.Context) {
			err := videoController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(200, gin.H{"ok": "ok"})
			}
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	server.Run(":8080")
}
