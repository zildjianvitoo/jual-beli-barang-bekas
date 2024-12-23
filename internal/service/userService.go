package service

import (
	"errors"
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/helper"
	"jual-beli-barang-bekas/internal/repository"
	"time"
)

type UserService struct {
	Repo   repository.UserRepository
	Auth   helper.Auth
	Config config.AppConfig
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

func (s UserService) isVerifiedUser(id uint) bool {

	currentUser, err := s.Repo.GetUserById(id)

	return err == nil && currentUser.Verified
}

func (s UserService) GetVerificationCode(e domain.User) (int, error) {
	if s.isVerifiedUser(e.ID) {
		return 0, errors.New("user already verified")
	}

	code, err := s.Auth.GenerateCode()
	if err != nil {
		return 0, err
	}

	// update user
	user := domain.User{
		Expiry: time.Now().Add(30 * time.Minute),
		Code:   code,
	}

	_, err = s.Repo.UpdateUser(e.ID, user)

	if err != nil {
		return 0, errors.New("unable to update verification code")
	}

	user, _ = s.Repo.GetUserById(e.ID)

	// send SMS
	// notificationClient := notification.NewNotificationClient(s.Config)

	// msg := fmt.Sprintf("Your verification code is %v", code)

	// err = notificationClient.SendSMS(user.Phone, msg)
	if err != nil {
		return 0, errors.New("error on sending sms")
	}

	return code, nil
}

func (s UserService) DoVerify(id uint, code int) error {
	if s.isVerifiedUser(id) {
		return errors.New("user already verified")
	}

	user, err := s.Repo.GetUserById(id)

	if err != nil {
		return err
	}

	if user.Code != code {
		return errors.New("verification code does not match")
	}

	if !time.Now().Before(user.Expiry) {
		return errors.New("verification code expired")
	}

	updateUser := domain.User{
		Verified: true,
	}

	_, err = s.Repo.UpdateUser(id, updateUser)

	if err != nil {
		return errors.New("unable to verify user")
	}

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
