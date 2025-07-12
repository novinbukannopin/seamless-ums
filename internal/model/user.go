package model

import (
	"github.com/go-playground/validator/v10"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Username    string    `json:"username" gorm:"column:username;unique;not null; type:varchar(50) ;validate:required"`
	Email       string    `json:"email" gorm:"column:email;not null; type:varchar(100) ;validate:required"`
	PhoneNumber string    `json:"phone_number" gorm:"column:phone_number;not null; type:varchar(15) ;validate:required"`
	Address     string    `json:"address" gorm:"column:address;not null; type:varchar(255) ;validate:required"`
	Dob         string    `json:"dob" gorm:"column:dob;not null; type:date"`
	Password    string    `json:"password,omitempty" gorm:"column:password;not null; type:varchar(255)"`
	FullName    string    `json:"full_name" gorm:"column:full_name;not null; type:varchar(100)"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}

func (User) TableName() string {
	return "users"
}

func (l User) Validate() error {
	v := validator.New()
	return v.Struct(l)
}

type UserSession struct {
	ID                  int `gorm:"primaryKey"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	UserID              int       `json:"user_id" gorm:"type:int; validate:required"`
	Token               string    `json:"token" gorm:"type:text; validate:required"`
	RefreshToken        string    `json:"refresh_token" gorm:"type:text; validate:required"`
	TokenExpired        time.Time `json:"-" gorm:"validate:required"`
	RefreshTokenExpired time.Time `json:"-" gorm:"validate:required"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}

func (l UserSession) Validate() error {
	v := validator.New()
	return v.Struct(l)
}
