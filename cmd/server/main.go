package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"payment-platform/config"
	"payment-platform/internal/container"
	"payment-platform/internal/http"
)

func main() {
	envFile := flag.String("env-file", ".env", ".env configuration file path")

	logrus.SetFormatter(&logrus.JSONFormatter{})
	configs := config.NewConfig(*envFile)
	router := gin.Default()

	cont := container.NewContainer(configs)
	http.SetupRoutes(router, cont)
	err := router.Run(":8080")
	if err != nil {
		fmt.Println(err.Error())
	}
}
