package controller

import (
	"net/http"
	"zregistry_service/model"
	handler "zregistry_service/request"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Register(ctx *gin.Context) {
	var register model.Register
	if err := handler.Binder(ctx, &register, c.log); err.Status != 0 {
		ctx.JSON(err.Status, err)

		return
	}

	err := c.service.Register(ctx, register)
	if err.Status != 0 {
		ctx.JSON(err.Status, err)

		return
	}

	ctx.JSON(http.StatusOK, handler.ErrResp{Status: http.StatusOK, Msg: "Registed Successfully."})
}
