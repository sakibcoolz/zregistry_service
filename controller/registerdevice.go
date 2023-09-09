package controller

import (
	"net/http"
	"zregistry_service/model"
	handler "zregistry_service/request"

	"github.com/gin-gonic/gin"
)

func (c *Controller) RegisterDevice(ctx *gin.Context) {
	var deviceinfo model.DeviceInfo

	if err := handler.Binder(ctx, &deviceinfo, c.log); err.Status != 0 {
		ctx.JSON(err.Status, err)

		return
	}

	token, err := c.service.RegisterDevice(ctx, deviceinfo)
	if err.Status != 0 {
		ctx.JSON(err.Status, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"device": token,
	})
}
