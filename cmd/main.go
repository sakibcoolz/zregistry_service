package main

import (
	"fmt"
	"zregistry_service/app"
	"zregistry_service/config"
	"zregistry_service/log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := log.NewLogger()

	config := config.NewConfig(logger)

	router := gin.Default()

	router = app.Apps(config, logger, router)

	if err := router.Run(fmt.Sprintf(":%s", config.GetService().Port)); err != nil {
		logger.Fatal("Service Error", zap.Error(err))
	}
}
