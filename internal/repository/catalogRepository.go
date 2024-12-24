package repository

import (
	"errors"
	"jual-beli-barang-bekas/internal/domain"
	"log"

	"gorm.io/gorm"
)

type CatalogRepository interface {
	CreateCategory(e *domain.Category) error
	GetCategories() ([]*domain.Category, error)
	GetCategoryById(id int) (*domain.Category, error)
	EditCategory(e *domain.Category) (*domain.Category, error)
	DeleteCategory(id int) error

	CreateProduct(e *domain.Product) error
	GetProducts() ([]*domain.Product, error)
	GetProductById(id int) (*domain.Product, error)
	GetSellerProducts(id int) ([]*domain.Product, error)
	EditProduct(e *domain.Product) (*domain.Product, error)
	DeleteProduct(e *domain.Product) error
}

type catalogRepository struct {
	db *gorm.DB
}

func NewCatalogRepository(db *gorm.DB) CatalogRepository {
	return &catalogRepository{
		db: db,
	}
}

func (c catalogRepository) CreateProduct(e *domain.Product) error {
	err := c.db.Model(&domain.Product{}).Create(e).Error

	if err != nil {
		log.Printf("err: %v", err)
		return errors.New("cannot create product")
	}

	return nil
}

func (c catalogRepository) GetProducts() ([]*domain.Product, error) {
	var products []*domain.Product

	err := c.db.Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c catalogRepository) GetProductById(id int) (*domain.Product, error) {
	var product *domain.Product

	err := c.db.First(&product, id).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("product does not exist")
	}

	return product, nil
}

func (c catalogRepository) GetSellerProducts(id int) ([]*domain.Product, error) {
	var products []*domain.Product

	err := c.db.Where("user_id=?", id).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (c catalogRepository) EditProduct(e *domain.Product) (*domain.Product, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("fail to update product")
	}

	return e, nil
}

func (c catalogRepository) DeleteProduct(e *domain.Product) error {
	err := c.db.Delete(&domain.Product{}, e.ID).Error

	if err != nil {
		return errors.New("product cannot delete")
	}

	return nil
}

func (c catalogRepository) CreateCategory(e *domain.Category) error {
	err := c.db.Create(&e).Error
	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("create category failed")
	}
	return nil
}

func (c catalogRepository) GetCategories() ([]*domain.Category, error) {
	var categories []*domain.Category

	err := c.db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c catalogRepository) GetCategoryById(id int) (*domain.Category, error) {
	var category *domain.Category

	err := c.db.First(&category, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("category does not exist")
	}

	return category, nil
}

func (c catalogRepository) EditCategory(e *domain.Category) (*domain.Category, error) {
	err := c.db.Save(&e).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return nil, errors.New("fail to update category")
	}

	return e, nil
}

func (c catalogRepository) DeleteCategory(id int) error {

	err := c.db.Delete(&domain.Category{}, id).Error

	if err != nil {
		log.Printf("db_err: %v", err)
		return errors.New("fail to delete category")
	}

	return nil

}
