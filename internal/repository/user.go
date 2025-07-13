package repository

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"seamless-ums/internal/model"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertNewUser(ctx context.Context, user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (model.User, error) {
	var (
		user model.User
		err  error
	)

	err = r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, gorm.ErrRecordNotFound
	}

	return user, nil
}

func (r *UserRepository) InsertNewUserSession(ctx context.Context, userSession *model.UserSession) error {
	return r.DB.Create(userSession).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, token string) error {
	return r.DB.Exec("DELETE FROM user_sessions WHERE token = ?", token).Error
}

func (r *UserRepository) UpdateTokenByRefreshToken(ctx context.Context, token string, refreshToken string) error {
	return r.DB.Exec("UPDATE user_sessions SET token = ? WHERE refresh_token = ?", token, refreshToken).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (model.UserSession, error) {
	var (
		userSession model.UserSession
		err         error
	)

	err = r.DB.Where("token = ?", token).First(&userSession).Error
	if err != nil {
		return userSession, err
	}

	if userSession.ID == 0 {
		return userSession, errors.Wrap(err, "user session not found")
	}

	return userSession, nil
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (model.UserSession, error) {
	var (
		userSession model.UserSession
		err         error
	)

	err = r.DB.Where("refresh_token = ?", refreshToken).First(&userSession).Error
	if err != nil {
		return userSession, err
	}

	if userSession.ID == 0 {
		return userSession, errors.Wrap(err, "user session not found")
	}

	return userSession, nil
}
