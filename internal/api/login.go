package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seamless-ums/constant"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
	"seamless-ums/internal/model"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log = helpers.Logger
		req model.LoginRequest
		res model.LoginResponse
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Info("Failed to parse JSON request body: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
	}

	if err := req.Validate(); err != nil {
		log.Info("Failed to validate request: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusBadRequest, constant.ErrFailedBadRequest, nil)
		return
	}

	res, err := api.LoginService.Login(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to login user: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusInternalServerError, constant.ErrFailedInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constant.SuccessMessage, res)
}
