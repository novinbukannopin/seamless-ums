package services

import (
	"context"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
	"seamless-ums/helpers"
	"seamless-ums/internal/interfaces"
	"seamless-ums/internal/model"
	"strconv"
	"time"
)

type LoginService struct {
	UserRepository interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req model.LoginRequest) (model.LoginResponse, error) {

	var (
		res model.LoginResponse
		now = time.Now()
	)

	userDetail, err := s.UserRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return res, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return res, errors.Wrap(err, "password mismatch")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "token", userDetail.Email, now)
	if err != nil {
		return res, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt_token", userDetail.Email, now)
	if err != nil {
		return res, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &model.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken["token"]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken["refresh_token"]),
	}

	err = s.UserRepository.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return res, errors.Wrap(err, "failed to insert new user session")
	}

	res.UserId = strconv.Itoa(userDetail.ID)
	res.Username = userDetail.Username
	res.FullName = userDetail.FullName
	res.Email = userDetail.Email
	res.Token = token
	res.RefreshToken = refreshToken
	return res, nil
}
