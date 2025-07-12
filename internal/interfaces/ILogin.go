package interfaces

import (
	"context"
	"seamless-ums/internal/model"
)

type ILoginService interface {
	Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error)
}
