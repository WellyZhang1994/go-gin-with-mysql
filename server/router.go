package server

import (
	"github.com/gin-gonic/gin"
	"github.com/WellyZhang1994/go-gin-with-mysql/controllers"
	"github.com/WellyZhang1994/go-gin-with-mysql/middlewares"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)


	router.Use(middlewares.AuthMiddleware())
	return router

}
