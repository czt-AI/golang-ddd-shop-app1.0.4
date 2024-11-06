package service

import (
	"context"
	"shop-app/model"
	"shop-app/repo"
)

// StatisticsService provides functionality to gather and report statistics.
type StatisticsService struct {
	repo StatisticsRepository
}

// GetUserStatistics retrieves user statistics.
func (s *StatisticsService) GetUserStatistics(ctx context.Context) ([]*model.UserStatistics, error) {
	return s.repo.GetUserStatistics(ctx)
}

// GetProductStatistics retrieves product statistics.
func (s *StatisticsService) GetProductStatistics(ctx context.Context) ([]*model.ProductStatistics, error) {
	return s.repo.GetProductStatistics(ctx)
}

// GetOrderStatistics retrieves order statistics.
func (s *StatisticsService) GetOrderStatistics(ctx context.Context) ([]*model.OrderStatistics, error) {
	return s.repo.GetOrderStatistics(ctx)
}

// GetPaymentStatistics retrieves payment statistics.
func (s *StatisticsService) GetPaymentStatistics(ctx context.Context) ([]*model.PaymentStatistics, error) {
	return s.repo.GetPaymentStatistics(ctx)
}

// GetLogisticsStatistics retrieves logistics statistics.
func (s *StatisticsService) GetLogisticsStatistics(ctx context.Context) ([]*model.LogisticsStatistics, error) {
	return s.repo.GetLogisticsStatistics(ctx)
}