package repository

import (
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	CreatePayment(payment *domain.Payment) error
	GetOrders(uId uint) ([]domain.OrderItem, error)
	GetOrderById(uId uint, id uint) (dto.SellerOrderDetails, error)
}

type transactionStorage struct {
	db *gorm.DB
}

func (t transactionStorage) CreatePayment(payment *domain.Payment) error {
	//TODO implement me
	panic("implement me")
}

func (t transactionStorage) GetOrders(uId uint) ([]domain.OrderItem, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionStorage) GetOrderById(uId uint, id uint) (dto.SellerOrderDetails, error) {
	//TODO implement
	panic("implement me")
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionStorage{db: db}
}
