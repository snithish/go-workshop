package main

import "github.com/gin-gonic/gin"

func InitializeRoutes() *gin.Engine {
	var router = gin.Default()
	return router
}
