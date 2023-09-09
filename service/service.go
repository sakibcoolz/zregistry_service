package service

import (
	"zregistry_service/domain"
	"zregistry_service/model"
	handler "zregistry_service/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service struct {
	log   *zap.Logger
	store domain.IStore
}

type IService interface {
	Register(ctx *gin.Context, register model.Register) handler.ErrResp
	Login(ctx *gin.Context, login model.Login) (string, handler.ErrResp)
	RegisterDevice(ctx *gin.Context, deviceinfo model.DeviceInfo) (string, handler.ErrResp)
}

func NewService(logger *zap.Logger, store domain.IStore) *Service {
	return &Service{
		log:   logger,
		store: store,
	}
}
