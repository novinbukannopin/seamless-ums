package api

import (
	"context"
	"fmt"
	"seamless-ums/cmd/proto/token_validation"
	"seamless-ums/constant"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
	"strconv"
)

type TokenValidationHandler struct {
	TokenValidationService interfaces.ITokenValidationService
	tokenvalidation.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenvalidation.TokenRequest) (*tokenvalidation.TokenResponse, error) {
	var (
		token = req.GetToken()
		log   = helpers.Logger
	)

	if token == "" {
		log.Error("Token is empty")
		err := fmt.Errorf("token is empty")
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	claimToken, err := s.TokenValidationService.TokenValidation(ctx, token)
	if err != nil {
		log.Error("Token validation failed", "error", err)
		return &tokenvalidation.TokenResponse{
			Message: err.Error(),
		}, nil
	}

	return &tokenvalidation.TokenResponse{
		Message: constant.SuccessMessage,
		Data: &tokenvalidation.UserData{
			UserId:   strconv.Itoa(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.Fullname,
		},
	}, nil
}
