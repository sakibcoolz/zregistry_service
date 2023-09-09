package domain

import (
	"zregistry_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Store) GetTenant(ctx *gin.Context, id uint) (model.TenantMaster, error) {
	var (
		tenant model.TenantMaster
		err    error
	)

	if err = s.db.First(&tenant, id).Error; err != nil {
		s.log.Error("failed to fetch tenant", zap.Error(err))

		return tenant, err
	}

	return tenant, err
}
