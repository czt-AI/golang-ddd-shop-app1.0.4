package performance_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"shop-app/handlers"
	"shop-app/service"
)

func TestAPIPerformance(t *testing.T) {
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

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.Handle("/users/", handlers.NewUserServiceHandler(userService).GetUser)
	server := httptest.NewServer(mux)

	// Start the timer
	startTime := time.Now()

	// Perform multiple requests
	for i := 0; i < 1000; i++ {
		resp, err := http.Get(server.URL + "/users/1")
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()
	}

	// Stop the timer
	duration := time.Since(startTime)

	// Calculate the average response time
	averageResponseTime := duration.Seconds() / 1000

	// Check if the average response time is within acceptable limits
	if averageResponseTime > 0.5 {
		t.Errorf("Expected average response time to be less than 0.5 seconds, got %f seconds", averageResponseTime)
	}

	// Check the server status
	if server.Error != nil {
		t.Errorf("Server encountered an error: %v", server.Error)
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