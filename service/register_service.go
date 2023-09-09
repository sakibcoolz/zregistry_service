package service

import (
	"errors"
	"time"
	"zregistry_service/model"
	handler "zregistry_service/request"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Service) Register(ctx *gin.Context, register model.Register) handler.ErrResp {
	var erros handler.ErrResp
	// Define the number of days to add
	daysToAdd := 3

	if err := s.tenantValidation(ctx, register.Tenent); err != nil {
		return handler.ValidationUnprocessableEntity(ctx, err, s.log)
	}

	if err := s.userValidation(ctx, register.User); err != nil {
		return handler.ValidationUnprocessableEntity(ctx, err, s.log)
	}

	register.Tenent.ExpiryDate = time.Now().Add(time.Duration(daysToAdd) * 24 * time.Hour)

	register.Tenent.Status = "ACTIVE"

	register.User.ExpiryDate = time.Now().Add(time.Duration(daysToAdd) * 24 * time.Hour)

	if err := s.store.Registry(ctx, register.Tenent, register.User); err != nil {
		return handler.InternalError(ctx, err, s.log)
	}

	return erros
}

func (s *Service) tenantValidation(ctx *gin.Context, tenent model.TenantMaster) error {
	if tenent.Name == "" || tenent.Domain == "" {
		err := errors.New("incompleted data for TenantMaster")
		s.log.Error("Validations error at service", zap.Any("ctx", ctx), zap.Error(err))

		return err
	}

	return nil
}

func (s *Service) userValidation(ctx *gin.Context, user model.UserMaster) error {
	if user.Name == "" ||
		user.Username == "" ||
		user.Contacts == nil ||
		user.UserType == "" ||
		user.Address == "" ||
		user.Password == "" {
		err := errors.New("incomplete data for UserMaster")
		s.log.Error("Validations error at service", zap.Error(err))

		return err
	}

	return nil
}
