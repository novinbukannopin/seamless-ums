package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"seamless-ums/internal/model"
)

type IRegisterService interface {
	Register(ctx context.Context, request model.User) (interface{}, error)
}

type IRegisterHandler interface {
	Register(c *gin.Context)
}
