package domain

import (
	"zregistry_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Store) Registry(ctx *gin.Context, tenent model.TenantMaster,
	users model.UserMaster) error {
	var err error

	if err = s.db.Save(&tenent).Error; err != nil {
		s.log.Error("Failed to store data")
		err = s.db.Model(&model.TenantMaster{}).Where("domain = ?", tenent.Domain).Find(&tenent).Error
		if err != nil {
			s.log.Error("Failed to get from TenantMaster", zap.Error(err))

			return err
		}
	}

	users.TenantMasterID = tenent.ID

	if err = s.db.Save(&users).Error; err != nil {
		s.log.Error("Failed to store user")

		return err
	}

	return err
}
