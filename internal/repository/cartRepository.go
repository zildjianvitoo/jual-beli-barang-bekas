package repository

import (
	"jual-beli-barang-bekas/internal/domain"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartRepository interface {
	GetCartItems(uId uint) ([]domain.Cart, error)
	GetCartItem(uId uint, pId uint) (domain.Cart, error)
	CreateCart(c domain.Cart) error
	UpdateCart(c domain.Cart) error
	DeleteCartById(id uint) error
	DeleteCartItems(uId uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r cartRepository) GetCartItems(uId uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id=?", uId).Find(&carts).Error
	return carts, err
}

func (r cartRepository) GetCartItem(uId uint, pId uint) (domain.Cart, error) {
	cartItem := domain.Cart{}
	err := r.db.Where("user_id=? AND product_id=?", uId, pId).First(&cartItem).Error
	return cartItem, err
}

func (r cartRepository) CreateCart(c domain.Cart) error {
	return r.db.Create(&c).Error
}

func (r cartRepository) UpdateCart(c domain.Cart) error {
	var cart domain.Cart
	err := r.db.Model(&cart).Clauses(clause.Returning{}).Where("id=?", c.ID).Updates(c).Error
	return err
}

func (r cartRepository) DeleteCartById(id uint) error {
	err := r.db.Delete(&domain.Cart{}, id).Error
	return err
}

func (r cartRepository) DeleteCartItems(uId uint) error {
	err := r.db.Where("user_id=?", uId).Delete(&domain.Cart{}).Error
	return err
}
