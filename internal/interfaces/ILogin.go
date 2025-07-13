package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"seamless-ums/internal/model"
)

type ILoginService interface {
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
}

type ILoginHandler interface {
	Login(c *gin.Context)
}
