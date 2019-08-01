package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	configureLogging()
	routes := InitializeRoutes()
	err := routes.Run()
	if err != nil {
		log.Error("Im bailing because ", err)
	}
}

func configureLogging() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}
