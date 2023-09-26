package internal

import (
	"context"
	"errors"
)

type ProductService interface {
	GetAll() *[]Product
	GetById(ctx context.Context) (*Product, error)
	Add(newProduct Product) *Product
	UpdateProduct(ctx context.Context, newProduct Product) *Product
	Delete(ctx context.Context) error
}

type service struct {
	storage ProductStorage
}

func NewService(storage ProductStorage) ProductService {
	return &service{storage: storage}
}

func (s *service) Add(newProduct Product) *Product {
	return s.storage.Add(newProduct)
}

func (s *service) Delete(ctx context.Context) error {
	rowProductID := ctx.Value("id")

	if rowProductID == nil {
		return errors.New("empty product id")
	}

	productId := rowProductID.(string)

	return s.storage.DeleteById(productId)
}

func (s *service) GetById(ctx context.Context) (*Product, error) {
	rowProductID := ctx.Value("id")
	if rowProductID == nil {
		return nil, errors.New("empty product id")
	}

	productId := rowProductID.(string)

	return s.storage.GetById(productId), nil
}

func (s *service) GetAll() *[]Product {
	return s.storage.GetAll()
}

func (s *service) UpdateProduct(ctx context.Context, newProduct Product) *Product {
	rowProductID := ctx.Value("id")

	productId := rowProductID.(string)

	return s.storage.Update(productId, newProduct)
}
