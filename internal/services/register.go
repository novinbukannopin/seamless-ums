package services

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"seamless-ums/internal/interfaces"
	"seamless-ums/internal/model"
)

type RegisterService struct {
	UserRepository interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, request model.User) (interface{}, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashPassword)
	err = s.UserRepository.InsertNewUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	response := request
	response.Password = ""
	return response, nil
}
