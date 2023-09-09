package domain

import (
	"zregistry_service/model"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Store) Login(ctx *gin.Context, login model.Login) (model.UserMaster, error) {
	var (
		err      error
		user     model.UserMaster
		contacts []model.Contact
	)

	err = s.db.Model(&model.UserMaster{}).
		Where("username = ? and password = ?", login.Username, login.Password).
		Scan(&user).Error
	if err != nil {
		s.log.Error("Failed To fetch login details", zap.Error(err), zap.Any("ctx", ctx))

		return user, err
	}

	err = s.db.Model(&model.Contact{}).
		Where("user_master_id = ?", user.ID).
		Scan(&contacts).Error
	if err != nil {
		s.log.Error("Failed To fetch contacts details", zap.Error(err), zap.Any("ctx", ctx))

		return user, err
	}

	user.Contacts = append(user.Contacts, contacts...)

	return user, err
}
