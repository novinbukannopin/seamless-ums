package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seamless-ums/constant"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
)

type LogoutHandler struct {
	LogoutService interfaces.ILogoutService
}

func (api *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	token := c.Request.Header.Get("Authorization")
	err := api.LogoutService.Logout(c.Request.Context(), token)
	if err != nil {
		log.Error("Logout failed", "error", err)
		helpers.SendResponseHTTP(
			c, http.StatusInternalServerError, constant.ErrFailedInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(
		c, http.StatusOK, constant.SuccessMessage, nil)
}
