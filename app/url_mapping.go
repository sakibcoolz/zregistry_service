package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"zregistry_service/config"
	"zregistry_service/controller"
	"zregistry_service/database"
	"zregistry_service/domain"
	"zregistry_service/middleware"
	"zregistry_service/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	StopService = make(chan os.Signal, 1)
)

func Apps(config *config.Configuration, logger *zap.Logger, router *gin.Engine) *gin.Engine {
	go TerminateService(StopService, logger)

	signal.Notify(StopService, syscall.SIGINT, syscall.SIGTERM)

	db := database.Connection(config.GetDatabase())

	if err := database.Migrations(db.DB); err != nil {
		logger.Fatal("Failed in DB Migrations")
	}

	store := domain.NewStore(logger, db.DB)

	service := service.NewService(logger, store)

	controller := controller.NewController(logger, service)

	router.Use(gin.Recovery())

	preapproute := router.Group(fmt.Sprintf("/%s", config.GetService().Name))

	preapproute.POST("/register", controller.Register)

	preapproute.POST("/login", controller.Login)

	approuter := router.Group("/apps")

	approuter.Use(middleware.JwtMiddlewares(logger))

	approuter.POST("/setdevice", controller.RegisterDevice)

	return router
}

func TerminateService(stopService chan os.Signal, log *zap.Logger) {
	log.Info("Service Started")
	select {
	case <-stopService:
		log.Info("Terminating service")

		os.Exit(0)
	}
}
