package interfaces

import (
	"context"
	"seamless-ums/internal/model"
)

type IUserRepository interface {
	InsertNewUser(ctx context.Context, user *model.User) error
	GetUserByUsername(ctx context.Context, username string) (model.User, error)
	InsertNewUserSession(ctx context.Context, userSession *model.UserSession) error
	DeleteUserSession(ctx context.Context, token string) error
	GetUserSessionByToken(ctx context.Context, token string) (model.UserSession, error)
	UpdateTokenByRefreshToken(ctx context.Context, token string, refreshToken string) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (model.UserSession, error)
}
