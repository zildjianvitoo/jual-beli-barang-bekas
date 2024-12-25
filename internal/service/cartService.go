package service

import (
	"errors"
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/helper"
	"jual-beli-barang-bekas/internal/repository"
	"log"
)

type CartService struct {
	Repo        repository.CartRepository
	CatalogRepo repository.CatalogRepository
	Auth        helper.Auth
	Config      config.AppConfig
}

func (s CartService) GetCart(id uint) ([]domain.Cart, error) {

	cartItems, err := s.Repo.GetCartItems(id)
	log.Printf("error %v", err)

	return cartItems, err
}

func (s CartService) AddItemToCart(input any, u domain.User) ([]*domain.Cart, error) {
	return nil, nil
}

func (s CartService) CreateCart(input dto.CreateCartRequest, u domain.User) ([]domain.Cart, error) {
	// Check if the cart is Exist
	cart, _ := s.Repo.GetCartItem(u.ID, input.ProductId)

	if cart.ID > 0 {
		if input.ProductId == 0 {
			return nil, errors.New("please provide a valid product id")
		}
		// Delete the cart item
		if input.Qty < 1 {
			err := s.Repo.DeleteCartById(cart.ID)
			if err != nil {
				log.Printf("Error on deleting cart item %v", err)
				return nil, errors.New("error on deleting cart item")
			}
		} else {
			// Update the cart item
			cart.Qty = input.Qty
			err := s.Repo.UpdateCart(cart)
			if err != nil {
				// log error
				return nil, errors.New("error on updating cart item")
			}
		}

	} else {

		product, _ := s.CatalogRepo.GetProductById(int(input.ProductId))
		if product.ID < 1 {
			return nil, errors.New("product not found to create cart item")
		}

		err := s.Repo.CreateCart(domain.Cart{
			UserId:    u.ID,
			ProductId: input.ProductId,
			Name:      product.Name,
			ImageUrl:  product.ImageUrl,
			Qty:       input.Qty,
			Price:     product.Price,
			SellerId:  uint(product.UserId),
		})

		if err != nil {
			return nil, errors.New("error on creating cart item")
		}
	}

	return s.Repo.GetCartItems(u.ID)

}
