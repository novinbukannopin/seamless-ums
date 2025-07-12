package interfaces

import (
	"context"
	"seamless-ums/internal/model"
)

type IRegisterService interface {
	Register(ctx context.Context, request model.User) (interface{}, error)
}
