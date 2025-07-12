package services

import (
	"context"
	"seamless-ums/internal/interfaces"
)

type LogoutService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	return s.UserRepository.DeleteUserSession(ctx, token)
}
