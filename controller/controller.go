package controller

import (
	"zregistry_service/service"

	"go.uber.org/zap"
)

type Controller struct {
	log     *zap.Logger
	service service.IService
}

func NewController(log *zap.Logger, service service.IService) *Controller {
	return &Controller{
		log:     log,
		service: service,
	}
}
