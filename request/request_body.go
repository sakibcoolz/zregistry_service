package handler

import (
	"net/http"
	"zregistry_service/literals"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
)

type ErrResp struct {
	Status int    `json:"status,omitempty"`
	Msg    string `json:"msg,omitempty"`
}

func Binder(ctx *gin.Context, obj any, log *zap.Logger) ErrResp {
	var errs ErrResp
	if err := ctx.ShouldBindBodyWith(obj, binding.JSON); err != nil {
		return ValidationError(ctx, err, log)
	}

	return errs
}

func ValidationError(ctx *gin.Context, err error, log *zap.Logger) ErrResp {
	errs := ErrResp{
		Status: http.StatusBadRequest,
		Msg:    literals.ErrorResponse[literals.REQUEST_VALIDATION_ERROR],
	}

	log.Error("Error in Body Validation", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}

func ValidationUnprocessableEntity(ctx *gin.Context, err error, log *zap.Logger) ErrResp {
	errs := ErrResp{
		Status: http.StatusUnprocessableEntity,
		Msg:    literals.ErrorResponse[literals.REQUEST_VALIDATION_ERROR],
	}

	log.Error("Error in Body Validation", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}

func InternalError(ctx *gin.Context, err error, log *zap.Logger) ErrResp {
	errs := ErrResp{
		Status: http.StatusInternalServerError,
		Msg:    literals.ErrorResponse[literals.FAIL_TO_STORE],
	}

	log.Error("Error in store", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}

func LoginFailed(ctx *gin.Context, err error, log *zap.Logger) ErrResp {
	errs := ErrResp{
		Status: http.StatusNoContent,
		Msg:    literals.ErrorResponse[literals.LOGIN_FAILED],
	}

	log.Error("Error in Body Validation", zap.Error(err), zap.Any("ctx", ctx), zap.Any("response", errs))

	return errs
}
