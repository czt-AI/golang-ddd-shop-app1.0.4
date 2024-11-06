package service

import (
	"context"
	"shop-app/model"
	"shop-app/repo"
)

// UserService provides functionality to manage users.
type UserService struct {
	userRepo repo.UserRepository
}

// GetUser retrieves a user by ID.
func (s *UserService) GetUser(ctx context.Context, id int) (*model.User, error) {
	return s.userRepo.GetUser(ctx, id)
}

// ListUsers retrieves a list of all users.
func (s *UserService) ListUsers(ctx context.Context) ([]*model.User, error) {
	return s.userRepo.ListUsers(ctx)
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

// UpdateUser updates an existing user.
func (s *UserService) UpdateUser(ctx context.Context, user *model.User) error {
	return s.userRepo.UpdateUser(ctx, user)
}

// DeleteUser deletes a user by ID.
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.userRepo.DeleteUser(ctx, id)
}