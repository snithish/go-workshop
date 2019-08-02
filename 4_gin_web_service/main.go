package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := InitializeRoutes()
	err := router.Run()
	if err != nil {
		logrus.Error("Im bailing because ", err)
	}
}
