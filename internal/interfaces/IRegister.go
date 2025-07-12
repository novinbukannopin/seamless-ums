package interfaces

import (
	"context"
	"seamless-ums/internal/model"
)

type RegisterRepository interface {
	InsertNewUser(ctx context.Context, user *model.User) error
}

type RegisterService interface {
	Register(ctx context.Context, request model.User) (interface{}, error)
}
