package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"seamless-ums/constant"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")
	if !ok {
		log.Error("Token not found in context")
		helpers.SendResponseHTTP(
			c, http.StatusUnauthorized, constant.ErrFailedInternalServer, nil)
		return
	}

	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("Failed to cast token claim")
		helpers.SendResponseHTTP(
			c, http.StatusUnauthorized, constant.ErrFailedInternalServer, nil)
		return
	}

	res, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)

	if err != nil {
		log.Error("RefreshToken failed", "error", err)
		helpers.SendResponseHTTP(
			c, http.StatusInternalServerError, constant.ErrFailedInternalServer, nil)
		return
	}

	helpers.SendResponseHTTP(
		c, http.StatusOK, constant.SuccessMessage, res)
}
