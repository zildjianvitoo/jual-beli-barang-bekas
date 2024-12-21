package service

import "jual-beli-barang-bekas/internal/domain"

type CartService struct {
}

func (s CartService) GetCart(userId uint) (*domain.Cart, error) {
	return nil, nil
}

func (s CartService) AddItemToCart(input any, u domain.User) ([]*domain.Cart, error) {
	return nil, nil
}
