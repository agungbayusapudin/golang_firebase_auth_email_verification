package service

import (
	"context"
	"crud_fire/model"
	"crud_fire/repository"
)

type ProducService interface {
	GetAllProducts(ctx context.Context) ([]model.Product, error)
	InsertProduct(ctx context.Context, product model.Product) error
	EditProduct(ctx context.Context, product model.Product) error
	DeleteProduct(ctx context.Context, product model.Product) error
}

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProducService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts(ctx context.Context) ([]model.Product, error) {
	product, err := s.repo.GetAllProducts(ctx)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *productService) InsertProduct(ctx context.Context, product model.Product) error {
	err := s.repo.InsertProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) EditProduct(ctx context.Context, product model.Product) error {
	err := s.repo.EditProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

func (s *productService) DeleteProduct(ctx context.Context, product model.Product) error {
	err := s.repo.DeleteProduct(ctx, product)
	if err != nil {
		return err
	}
	return nil
}
