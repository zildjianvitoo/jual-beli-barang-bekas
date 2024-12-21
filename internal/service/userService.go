package service

import "jual-beli-barang-bekas/internal/domain"

type UserService struct {
}

func (s UserService) Register(input any) (string, error) {

	return "", nil
}

func (s UserService) Login(input any) (string, error) {
	return "", nil
}

func (s UserService) GetUserByEmail(email string) (*domain.User, error) {
	// Business Logic

	return nil, nil
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	return 0, nil
}

func (s UserService) DoVerify(id uint, code uint) error {
	return nil
}

func (s UserService) GetProfile(id uint) (*domain.User, error) {
	return nil, nil
}

func (s UserService) CreateProfile(id uint, input any) error {
	return nil
}

func (s UserService) UpdateProfile(id uint, input any) error {
	return nil
}

func (s UserService) BecomeSeller(id uint, input any) (string, error) {
	return "", nil
}
