package service_test

import (
	"context"
	"testing"
	"shop-app/model"
	"shop-app/service"
)

func TestUserService_GetUser(t *testing.T) {
	// Setup test environment
	userService := service.NewUserService()
	user := &model.User{
		ID:        1,
		Username:  "testuser",
		Password:  "password",
		Email:     "test@example.com",
		Phone:     "1234567890",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Mock database repository
	userRepo := &mockUserRepository{user: user}
	userService.userRepo = userRepo

	// Call the method
	gotUser, err := userService.GetUser(context.Background(), 1)
	if err != nil {
		t.Errorf("UserService.GetUser() error = %v", err)
	}

	// Assert the result
	if gotUser.ID != user.ID {
		t.Errorf("UserService.GetUser() = %v, want %v", gotUser.ID, user.ID)
	}
	if gotUser.Username != user.Username {
		t.Errorf("UserService.GetUser() = %v, want %v", gotUser.Username, user.Username)
	}
	if gotUser.Email != user.Email {
		t.Errorf("UserService.GetUser() = %v, want %v", gotUser.Email, user.Email)
	}
}

type mockUserRepository struct {
	user *model.User
}

func (m *mockUserRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	return m.user, nil
}

func (m *mockUserRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	return nil, nil
}

func (m *mockUserRepository) CreateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (m *mockUserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	return nil
}

func (m *mockUserRepository) DeleteUser(ctx context.Context, id int) error {
	return nil
}