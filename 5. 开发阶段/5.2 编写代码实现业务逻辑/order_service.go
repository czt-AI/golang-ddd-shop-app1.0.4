package service

import (
	"context"
	"shop-app/model"
	"shop-app/repo"
)

// OrderService provides functionality to manage orders.
type OrderService struct {
	orderRepo repo.OrderRepository
	userRepo  repo.UserRepository
	productRepo repo.ProductRepository
}

// CreateOrder creates a new order.
func (s *OrderService) CreateOrder(ctx context.Context, order *model.Order) error {
	// Check product availability
	for _, prodId := range order.ProductIds {
		prod, err := s.productRepo.GetProduct(ctx, prodId)
		if err != nil {
			return err
		}
		if prod Stock < order.Quantity {
			return fmt.Errorf("product with ID %d is out of stock", prodId)
		}
	}

	// Deduct product stock
	for _, prodId := range order.ProductIds {
		err := s.productRepo.UpdateStock(ctx, prodId, prod.Stock-order.Quantity)
		if err != nil {
			return err
		}
	}

	// Create order
	return s.orderRepo.CreateOrder(ctx, order)
}

// ListOrders retrieves a list of all orders.
func (s *OrderService) ListOrders(ctx context.Context) ([]*model.Order, error) {
	return s.orderRepo.ListOrders(ctx)
}

// GetOrder retrieves an order by ID.
func (s *OrderService) GetOrder(ctx context.Context, id int) (*model.Order, error) {
	return s.orderRepo.GetOrder(ctx, id)
}

// UpdateOrder updates an existing order.
func (s *OrderService) UpdateOrder(ctx context.Context, order *model.Order) error {
	return s.orderRepo.UpdateOrder(ctx, order)
}

// DeleteOrder deletes an order by ID.
func (s *OrderService) DeleteOrder(ctx context.Context, id int) error {
	return s.orderRepo.DeleteOrder(ctx, id)
}