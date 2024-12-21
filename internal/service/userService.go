package service

import (
	"errors"
	"fmt"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s UserService) Register(input dto.UserRegister) (string, error) {
	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: input.Password,
		Phone:    input.Phone,
	})
	userData := fmt.Sprint(user.Email, " ", user.Phone)
	// generate token

	return userData, err
}

func (s UserService) Login(loginInput dto.UserLogin) (string, error) {
	user, err := s.Repo.GetUser(loginInput.Email)

	if err != nil {
		return "", errors.New("user does not exist")
	}

	return user.Email, nil
}

func (s UserService) getUserByEmail(email string) (*domain.User, error) {
	user, err := s.Repo.GetUser(email)

	return &user, err
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) DoVerify(id uint, code uint) error {
	return nil
}

func (s UserService) GetProfile(userId uint) (*domain.User, error) {

	return nil, nil
}

func (s UserService) CreateProfile(userId uint, input any) error {
	return nil
}

func (s UserService) UpdateProfile(userId uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(userId uint, input any) (string, error) {
	return "", nil
}
