package repository

import (
	"context"
	"gorm.io/gorm"
	"seamless-ums/internal/model"
)

type RegisterRepository struct {
	DB *gorm.DB
}

func (r *RegisterRepository) InsertNewUser(ctx context.Context, user *model.User) error {
	return r.DB.Create(user).Error
}
