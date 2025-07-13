package services

import (
	"context"
	"github.com/pkg/errors"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
	"seamless-ums/internal/model"
	"time"
)

type RefreshTokenService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (model.RefreshTokenResponse, error) {
	resp := model.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Fullname, "token", tokenClaim.Email, time.Now())
	if err != nil {
		return resp, errors.Wrap(err, "failed to generate new token")
	}

	err = s.UserRepository.UpdateTokenByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return resp, errors.Wrap(err, "failed to update token by refresh token")
	}

	resp.Token = token
	return resp, nil
}
