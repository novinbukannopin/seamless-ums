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
	)

	userDetail, err := s.UserRepository.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return res, errors.Wrap(err, "failed to get user by username")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password)); err != nil {
		return res, errors.Wrap(err, "password mismatch")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", userDetail.Email, time.Now())
	if err != nil {
		return res, errors.Wrap(err, "failed to generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.ID, userDetail.Username, userDetail.FullName, "jwt", userDetail.Email, time.Now())
	if err != nil {
		return res, errors.Wrap(err, "failed to generate refresh token")
	}

	userSession := &model.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        time.Now().Add(24 * time.Hour),
		RefreshTokenExpired: time.Now().Add(30 * 24 * time.Hour),
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
