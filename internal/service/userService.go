package service

import (
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"
	"log"
)

type UserService struct {
}

func (s UserService) Register(input dto.UserRegister) (string, error) {
	log.Println(input)
	return "tokenennaenamea", nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

func (s UserService) getUserByEmail(email string) (*domain.User, error) {
	// Business Logic

	return nil, nil
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
