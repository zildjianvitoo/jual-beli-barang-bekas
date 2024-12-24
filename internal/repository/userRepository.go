package repository

import (
	"errors"
	"jual-beli-barang-bekas/internal/domain"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepository interface {
	CreateUser(user domain.User) (domain.User, error)
	GetUser(email string) (domain.User, error)
	GetUserById(userId uint) (domain.User, error)
	UpdateUser(userId uint, user domain.User) (domain.User, error)
	CreateBankAccount(entity domain.BankAccount) error

	// TODO: selesaikan
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(user domain.User) (domain.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		log.Printf("Error when creating user %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return user, nil
}

func (r userRepository) GetUser(email string) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, "email=?", email).Error

	if err != nil {
		log.Printf("Error when getting user by email %v", err)
		return domain.User{}, errors.New("error when getting user by email")
	}

	return user, nil
}

func (r userRepository) GetUserById(userId uint) (domain.User, error) {
	var user domain.User

	err := r.db.First(&user, userId).Error

	if err != nil {
		log.Printf("Get user error %v", err)
		return domain.User{}, errors.New("error when getting user by id")
	}

	return user, nil
}
func (r userRepository) UpdateUser(userId uint, userInput domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", userId).Updates(userInput).Error

	if err != nil {
		log.Printf("Update user error %v", err)
		return domain.User{}, errors.New("error when updating user")
	}

	return user, nil
}

func (r userRepository) CreateBankAccount(entity domain.BankAccount) error {
	return r.db.Create(&entity).Error
}
