package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"seamless-ums/helpers"
	"seamless-ums/internal/model"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (model.RefreshTokenResponse, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(c *gin.Context)
}
