package main

import (
	"github.com/sivalabs-bookstore/promotion_service_go/src/application"
)

func main() {
	appConfig := application.AppConfig{
		Port: 3000,
	}
	cleanUpTasks := application.Setup(appConfig)
	defer cleanUpTasks()
}
