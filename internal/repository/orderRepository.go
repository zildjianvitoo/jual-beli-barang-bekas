package repository

import (
	"errors"
	"jual-beli-barang-bekas/internal/domain"
	"log"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(o domain.Order) error
	GetOrders(uId uint) ([]domain.Order, error)
	GetOrderById(id uint, uId uint) (domain.Order, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{
		db: db,
	}
}

func (r orderRepository) CreateOrder(o domain.Order) error {
	err := r.db.Create(&o).Error
	if err != nil {
		log.Printf("error on creating order %v", err)
		return errors.New("failed to create order")
	}
	return nil
}

func (r orderRepository) GetOrders(uId uint) ([]domain.Order, error) {
	var orders []domain.Order
	err := r.db.Where("user_id=?", uId).Find(&orders).Error

	if err != nil {
		log.Printf("error on fetching orders %v", err)
		return nil, errors.New("failed to fetch orders")
	}

	return orders, nil
}

func (r orderRepository) GetOrderById(id uint, uId uint) (domain.Order, error) {
	var order domain.Order
	err := r.db.Preload("Items").Where("id=? AND user_id=?", id, uId).First(&order).Error

	if err != nil {
		log.Printf("error on fetching order %v", err)
		return domain.Order{}, errors.New("failed to fetch order")
	}

	return order, nil
}
