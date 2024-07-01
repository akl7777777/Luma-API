package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"luma-api/docs"
	"luma-api/middleware"
)

func RegisterRouter(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiRouter := r.Group("/luma/v1", middleware.SecretAuth())
	{
		apiRouter.POST("/generations/", Generations)
		apiRouter.GET("/generations/*action", Fetch)
		apiRouter.GET("/task", Task)
		apiRouter.POST("/generations/file_upload", Upload)
	}
}
