package service

import (
	"context"
	"shop-app/model"
	"shop-app/repo"
)

// LogisticsService provides functionality to manage logistics information.
type LogisticsService struct {
	logisticsRepo repo.LogisticsRepository
}

// GetLogisticsInfo retrieves logistics information by order ID.
func (s *LogisticsService) GetLogisticsInfo(ctx context.Context, orderId int) (*model.Logistics, error) {
	return s.logisticsRepo.GetLogisticsInfo(ctx, orderId)
}

// UpdateLogisticsInfo updates logistics information.
func (s *LogisticsService) UpdateLogisticsInfo(ctx context.Context, logistics *model.Logistics) error {
	return s.logisticsRepo.UpdateLogisticsInfo(ctx, logistics)
}

// ListLogisticsInfo retrieves a list of all logistics information.
func (s *LogisticsService) ListLogisticsInfo(ctx context.Context) ([]*model.Logistics, error) {
	return s.logisticsRepo.ListLogisticsInfo(ctx)
}

// CreateLogisticsInfo creates new logistics information.
func (s *LogisticsService) CreateLogisticsInfo(ctx context.Context, logistics *model.Logistics) error {
	return s.logisticsRepo.CreateLogisticsInfo(ctx, logistics)
}

// DeleteLogisticsInfo deletes logistics information by ID.
func (s *LogisticsService) DeleteLogisticsInfo(ctx context.Context, logisticsId int) error {
	return s.logisticsRepo.DeleteLogisticsInfo(ctx, logisticsId)
}