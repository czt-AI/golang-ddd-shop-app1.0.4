package integration_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"shop-app/handlers"
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

	// Setup HTTP server
	mux := http.NewServeMux()
	mux.Handle("/users/", handlers.NewUserServiceHandler(userService).GetUser)
	server := httptest.NewServer(mux)

	// Test GET /users/1
	resp, err := http.Get(server.URL + "/users/1")
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	// Check the response body
	var userResp model.User
	if err := json.NewDecoder(resp.Body).Decode(&userResp); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if userResp.ID != user.ID {
		t.Errorf("Expected user ID %d, got %d", user.ID, userResp.ID)
	}
	if userResp.Username != user.Username {
		t.Errorf("Expected user username %s, got %s", user.Username, userResp.Username)
	}
	if userResp.Email != user.Email {
		t.Errorf("Expected user email %s, got %s", user.Email, userResp.Email)
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