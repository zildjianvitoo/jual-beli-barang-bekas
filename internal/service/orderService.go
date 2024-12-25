package service

import (
	"errors"
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/helper"
	"jual-beli-barang-bekas/internal/repository"
	"log"
)

type OrderService struct {
	Repo     repository.OrderRepository
	CartRepo repository.CartRepository
	Auth     helper.Auth
	Config   config.AppConfig
}

func (s OrderService) CreateOrder(u domain.User) (int, error) {

	cartItems, err := s.CartRepo.GetCartItems(u.ID)
	if err != nil {
		return 0, errors.New("error on finding cart items")
	}

	if len(cartItems) == 0 {
		return 0, errors.New("cart is empty cannot create the order")
	}

	paymentId := "PAY12345"
	txnId := "TXN12345"
	orderRef, _ := helper.RandomNumbers(8)

	var amount float64
	var orderItems []domain.OrderItem

	for _, item := range cartItems {
		amount += item.Price * float64(item.Qty)
		orderItems = append(orderItems, domain.OrderItem{
			ProductId: item.ProductId,
			Qty:       item.Qty,
			Price:     item.Price,
			Name:      item.Name,
			ImageUrl:  item.ImageUrl,
			SellerId:  item.SellerId,
		})
	}

	order := domain.Order{
		UserId:         u.ID,
		PaymentId:      paymentId,
		TransactionId:  txnId,
		OrderRefNumber: uint(orderRef),
		Amount:         amount,
		Items:          orderItems,
	}

	err = s.Repo.CreateOrder(order)
	if err != nil {
		return 0, err
	}

	err = s.CartRepo.DeleteCartItems(u.ID)
	log.Printf("Deleting cart items Error %v", err)

	return orderRef, nil
}

func (s OrderService) GetOrders(u domain.User) ([]domain.Order, error) {
	orders, err := s.Repo.GetOrders(u.ID)

	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s OrderService) GetOrderById(id uint, uId uint) (domain.Order, error) {
	order, err := s.Repo.GetOrderById(id, uId)

	if err != nil {
		return order, err
	}

	return order, nil
}
