package domain

import (
	"zregistry_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Store struct {
	log *zap.Logger
	db  *gorm.DB
}

type IStore interface {
	Registry(ctx *gin.Context, Tenent model.TenantMaster, users model.UserMaster) error
	Login(ctx *gin.Context, login model.Login) (model.UserMaster, error)
	GetTenant(ctx *gin.Context, id uint) (model.TenantMaster, error)
	RegisterDevice(ctx *gin.Context, deviceinfo model.DeviceInfo) error
}

func NewStore(logger *zap.Logger, db *gorm.DB) *Store {
	return &Store{
		log: logger,
		db:  db,
	}
}
