package services

import (
	"context"
	"github.com/pkg/errors"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
)

type TokenValidationService struct {
	UserRepository interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var (
		res *helpers.ClaimToken
		log = helpers.Logger
	)
	res, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		log.Error("Token validation failed", "error", err)
		return res, errors.Wrap(err, "Token validation failed")
	}

	_, err = s.UserRepository.GetUserSessionByToken(ctx, token)
	if err != nil {
		log.Error("Failed to get user session by token", "error", err)
		return res, errors.Wrap(err, "Failed to get user session by token")
	}

	return res, nil
}
