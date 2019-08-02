package main

import (
	"4_gin_web_service/controllers"
	"4_gin_web_service/repositories"
	"4_gin_web_service/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitializeRoutes() *gin.Engine {
	var router = gin.Default()
	//repositories
	petRepository := repositories.NewPetRepository()
	//Services
	petService := services.NewPetService(petRepository)
	//Controllers
	petController := controllers.NewPetController(petService)

	petGroup := router.Group("/pet")
	{
		petGroup.POST("", getGinHandler(petController.CreatePet))
		petGroup.PUT("", getGinHandler(petController.UpdatePet))
		petGroup.DELETE("/:petID", getGinHandler(petController.DeletePet))
		petGroup.GET("/:petID", getGinHandler(petController.GetPet))
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
