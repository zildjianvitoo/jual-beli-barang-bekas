package service

import (
	"errors"
	"jual-beli-barang-bekas/config"
	"jual-beli-barang-bekas/internal/domain"
	"jual-beli-barang-bekas/internal/dto"
	"jual-beli-barang-bekas/internal/helper"
	"jual-beli-barang-bekas/internal/repository"
)

type CatalogService struct {
	Repo   repository.CatalogRepository
	Auth   helper.Auth
	Config config.AppConfig
}

func (s CatalogService) CreateCategory(input dto.CreateCategoryRequest) error {

	err := s.Repo.CreateCategory(&domain.Category{
		Name:         input.Name,
		ImageUrl:     input.ImageUrl,
		DisplayOrder: input.DisplayOrder,
	})

	return err
}

func (s CatalogService) EditCategory(id int, input dto.CreateCategoryRequest) (*domain.Category, error) {

	existCat, err := s.Repo.GetCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")

	}

	if len(input.Name) > 0 {
		existCat.Name = input.Name
	}

	if input.ParentId > 0 {
		existCat.ParentId = input.ParentId
	}

	if len(input.ImageUrl) > 0 {
		existCat.ImageUrl = input.ImageUrl
	}

	if input.DisplayOrder > 0 {
		existCat.DisplayOrder = input.DisplayOrder
	}

	updatedCat, err := s.Repo.EditCategory(existCat)

	return updatedCat, err
}

func (s CatalogService) DeleteCategory(id int) error {
	err := s.Repo.DeleteCategory(id)
	if err != nil {
		return errors.New("category does not exist to delete")
	}

	return nil
}

func (s CatalogService) GetCategories() ([]*domain.Category, error) {

	categories, err := s.Repo.GetCategories()
	if err != nil {
		return nil, errors.New("categories does not exist")
	}

	return categories, err

}

func (s CatalogService) GetCategory(id int) (*domain.Category, error) {
	cat, err := s.Repo.GetCategoryById(id)
	if err != nil {
		return nil, errors.New("category does not exist")

	}
	return cat, nil
}

func (s CatalogService) CreateProduct(input dto.CreateProductRequest, user domain.User) error {
	err := s.Repo.CreateProduct(&domain.Product{
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryId:  input.CategoryId,
		ImageUrl:    input.ImageUrl,
		UserId:      int(user.ID),
	})

	return err
}

func (s CatalogService) EditProduct(id int, input dto.CreateProductRequest, user domain.User) (*domain.Product, error) {

	exitProduct, err := s.Repo.GetProductById(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	// verify product owner
	if exitProduct.UserId != int(user.ID) {
		return nil, errors.New("you dont have manage rights of this product")
	}

	if len(input.Name) > 0 {
		exitProduct.Name = input.Name
	}

	if len(input.Description) > 0 {
		exitProduct.Description = input.Description
	}

	if input.Price > 0 {
		exitProduct.Price = input.Price
	}

	if input.CategoryId > 0 {
		exitProduct.CategoryId = input.CategoryId
	}

	updatedProduct, err := s.Repo.EditProduct(exitProduct)

	return updatedProduct, err
}

func (s CatalogService) DeleteProduct(id int, user domain.User) error {
	exitProduct, err := s.Repo.GetProductById(id)
	if err != nil {
		return errors.New("product does not exist")
	}

	// verify product owner
	if exitProduct.UserId != int(user.ID) {
		return errors.New("you don't have manage rights of this product")
	}

	err = s.Repo.DeleteProduct(exitProduct)
	if err != nil {
		return errors.New("product cannot delete")
	}

	return nil
}

func (s CatalogService) GetProducts() ([]*domain.Product, error) {
	products, err := s.Repo.GetProducts()
	if err != nil {
		return nil, errors.New("products does not exist")
	}

	return products, err
}

func (s CatalogService) GetProductById(id int) (*domain.Product, error) {
	product, err := s.Repo.GetProductById(id)
	if err != nil {
		return nil, errors.New("product does not exist")
	}

	return product, nil
}

func (s CatalogService) GetSellerProducts(id int) ([]*domain.Product, error) {
	products, err := s.Repo.GetSellerProducts(id)
	if err != nil {
		return nil, errors.New("products does not exist")
	}

	return products, err
}
