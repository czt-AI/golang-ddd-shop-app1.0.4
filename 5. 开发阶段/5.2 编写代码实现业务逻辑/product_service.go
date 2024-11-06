package service

import (
	"context"
	"shop-app/model"
	"shop-app/repo"
)

// ProductService provides functionality to manage products.
type ProductService struct {
	productRepo repo.ProductRepository
}

// GetProduct retrieves a product by ID.
func (s *ProductService) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	return s.productRepo.GetProduct(ctx, id)
}

// ListProducts retrieves a list of all products.
func (s *ProductService) ListProducts(ctx context.Context) ([]*model.Product, error) {
	return s.productRepo.ListProducts(ctx)
}

// CreateProduct creates a new product.
func (s *ProductService) CreateProduct(ctx context.Context, product *model.Product) error {
	return s.productRepo.CreateProduct(ctx, product)
}

// UpdateProduct updates an existing product.
func (s *ProductService) UpdateProduct(ctx context.Context, product *model.Product) error {
	return s.productRepo.UpdateProduct(ctx, product)
}

// DeleteProduct deletes a product by ID.
func (s *ProductService) DeleteProduct(ctx context.Context, id int) error {
	return s.productRepo.DeleteProduct(ctx, id)
}