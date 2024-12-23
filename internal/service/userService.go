package service

import (
	"errors"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/helper"
	"jual-beli-barang-bekas/internal/repository"
)

type UserService struct {
	Repo repository.UserRepository
	Auth helper.Auth
}

func (s UserService) Register(input dto.UserRegister) (string, error) {
	hashedPassword, err := s.Auth.CreateHashedPassword(input.Password)

	if err != nil {
		return "", err
	}

	user, err := s.Repo.CreateUser(domain.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
}

func (s UserService) Login(loginInput dto.UserLogin) (string, error) {
	user, err := s.getUserByEmail(loginInput.Email)

	if err != nil {
		return "", errors.New("user does not exist")
	}

	err = s.Auth.VerifyPassword(loginInput.Password, user.Password)

	if err != nil {
		return "", err
	}

	return s.Auth.GenerateToken(user.ID, user.Email, user.UserType)
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
