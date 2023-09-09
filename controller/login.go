package controller

import (
	"net/http"
	"zregistry_service/model"
	handler "zregistry_service/request"

	"github.com/gin-gonic/gin"
)

func (c *Controller) Login(ctx *gin.Context) {
	var login model.Login

	if err := handler.Binder(ctx, &login, c.log); err.Status != 0 {
		ctx.JSON(err.Status, err)

		return
	}

	token, err := c.service.Login(ctx, login)
	if err.Status != 0 {
		ctx.JSON(err.Status, err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
