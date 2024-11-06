package service

import (
	"context"
	"shop-app/model"
	"shop-app/repo"
)

// PaymentService provides functionality to manage payments.
type PaymentService struct {
	paymentRepo repo.PaymentRepository
}

// CreatePayment creates a new payment record.
func (s *PaymentService) CreatePayment(ctx context.Context, payment *model.Payment) error {
	return s.paymentRepo.CreatePayment(ctx, payment)
}

// GetPayment retrieves a payment by ID.
func (s *PaymentService) GetPayment(ctx context.Context, id int) (*model.Payment, error) {
	return s.paymentRepo.GetPayment(ctx, id)
}

// ListPayments retrieves a list of all payments.
func (s *PaymentService) ListPayments(ctx context.Context) ([]*model.Payment, error) {
	return s.paymentRepo.ListPayments(ctx)
}

// UpdatePayment updates an existing payment record.
func (s *PaymentService) UpdatePayment(ctx context.Context, payment *model.Payment) error {
	return s.paymentRepo.UpdatePayment(ctx, payment)
}

// DeletePayment deletes a payment by ID.
func (s *PaymentService) DeletePayment(ctx context.Context, id int) error {
	return s.paymentRepo.DeletePayment(ctx, id)
}