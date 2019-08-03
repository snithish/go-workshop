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
		petGroup.POST("", petController.CreatePet)
		petGroup.PUT("", petController.UpdatePet)
		petGroup.DELETE("/:petID", petController.DeletePet)
		petGroup.GET("/:petID", petController.GetPet)
	}

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{})
	})
	return router
}
