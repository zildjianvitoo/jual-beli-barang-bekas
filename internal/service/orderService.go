package service

import "jual-beli-barang-bekas/internal/domain"

type OrderService struct{}

func (s UserService) CreateOrder(u domain.User) (int, error) {
	return 0, nil
}

func (s UserService) GetOrders(u domain.User) ([]any, error) {
	return nil, nil
}

func (s UserService) GetOrderById(orderId int, userId int) (*domain.Order, error) {
	return nil, nil
}
