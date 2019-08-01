package main

import (
	"4_gin_web_service/src/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRoutes() *gin.Engine {
	var router = gin.Default()
	petController := controllers.NewPetController()

	petGroup := router.Group("/pet")
	{
		petGroup.POST("", getGinHandler(petController.CreatePet))
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})
	return router
}

func getGinHandler(appHandler controllers.AppContextHandlerFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		appHandler(context)
	}
}
