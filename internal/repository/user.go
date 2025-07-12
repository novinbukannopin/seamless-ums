package repository

import (
	"context"
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
