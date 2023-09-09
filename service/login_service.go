package service

import (
	"errors"
	"zregistry_service/model"
	handler "zregistry_service/request"
	"zregistry_service/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Service) Login(ctx *gin.Context, login model.Login) (string, handler.ErrResp) {
	var (
		token   string
		errresp handler.ErrResp
	)

	if err := s.loginValidation(ctx, login); err != nil {
		errresp = handler.ValidationUnprocessableEntity(ctx, err, s.log)

		return "", errresp
	}

	user, err := s.store.Login(ctx, login)
	if err != nil {
		errresp = handler.LoginFailed(ctx, err, s.log)

		return "", errresp
	}

	tenant, err := s.store.GetTenant(ctx, user.TenantMasterID)
	if err != nil {
		errresp = handler.InternalError(ctx, err, s.log)

		return "", errresp
	}

	tenant.Users = append(tenant.Users, user)

	token, err = utils.GenerateToken(tenant)
	if err != nil {
		errresp = handler.InternalError(ctx, err, s.log)

		return "", errresp
	}

	return token, errresp
}

func (s *Service) loginValidation(ctx *gin.Context, login model.Login) error {
	var err error

	if login.Password == "" || login.Username == "" {
		err = errors.New("Invalid login")

		s.log.Info("Invalid login", zap.Any("ctx", ctx), zap.Error(err))
	}

	return err
}
