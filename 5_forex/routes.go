package main

import (
	"forex/controllers"
	"forex/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func InitializeRoutes() *gin.Engine {
	var router = gin.Default()
	conversionController := controllers.NewConversionController()
	router.POST("/v1/convert", conversionController.Convert)

	// Because of lack of generic body can't be generalized as in Spring
	// One approach is controller accepting type unsafe interface{} similar to Object in Java
	// Avoid at all costs
	router.POST("/v1/lamda/convert", func(context *gin.Context) {
		var request models.ConversionRequest
		if err := context.ShouldBindBodyWith(&request, binding.JSON); err != nil {
			context.JSON(400, "")
		}
		conversionController.ConvertClosureBody(request, context)
	})
	return router
}
