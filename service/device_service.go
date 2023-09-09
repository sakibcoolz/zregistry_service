package service

import (
	"errors"
	"fmt"
	"time"
	"zregistry_service/literals"
	"zregistry_service/model"
	handler "zregistry_service/request"
	"zregistry_service/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Service) RegisterDevice(ctx *gin.Context, deviceinfo model.DeviceInfo) (string, handler.ErrResp) {
	var (
		err handler.ErrResp
	)

	if err := s.deviceInfoValidation(ctx, deviceinfo); err != nil {
		return "", handler.InternalError(ctx, err, s.log)
	}

	str := utils.MD5Encryption(s.log, fmt.Sprintf("%d|%s", deviceinfo.UserMasterID, deviceinfo.Hostname))
	if str == "" {
		errs := errors.New("failed encryption")

		return "", handler.InternalError(ctx, errs, s.log)
	}

	deviceinfo.DeviceKey = str
	deviceinfo.Status = literals.ACTIVE
	deviceinfo.ExpiryDate = time.Now().Add(time.Duration(3) * 24 * time.Hour)

	if errs := s.store.RegisterDevice(ctx, deviceinfo); errs != nil {
		err = handler.InternalError(ctx, errs, s.log)

		return "", err
	}

	return str, err
}

func (s *Service) deviceInfoValidation(ctx *gin.Context, deviceinfo model.DeviceInfo) error {
	var err error

	if deviceinfo.UserMasterID == 0 || deviceinfo.Hostname == "" || deviceinfo.Platform == "" {
		err = errors.New("validation failed for device_info")
		s.log.Error("error on ", zap.Error(err))
	}

	return err
}
