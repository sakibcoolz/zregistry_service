package domain

import (
	"zregistry_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Store) RegisterDevice(ctx *gin.Context, deviceinfo model.DeviceInfo) error {
	var err error

	if err := s.db.Save(&deviceinfo).Error; err != nil {
		s.log.Error("failed to store deviceinfo", zap.Error(err), zap.Any("ctx", ctx))

		return err
	}

	return err
}
