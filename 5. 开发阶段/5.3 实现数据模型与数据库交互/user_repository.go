package repo

import (
	"database/sql"
	"shop-app/model"
)

// UserRepository provides functionality to manage user data in the database.
type UserRepository struct {
	db *sql.DB
}

// GetUser retrieves a user by ID.
func (ur *UserRepository) GetUser(ctx context.Context, id int) (*model.User, error) {
	// SQL query to retrieve user by ID
	query := `SELECT id, username, password, email, phone, created_at, updated_at FROM users WHERE id = ?`
	rows, err := ur.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user model.User
	if rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
	}
	return &user, nil
}

// ListUsers retrieves a list of all users.
func (ur *UserRepository) ListUsers(ctx context.Context) ([]*model.User, error) {
	// SQL query to retrieve all users
	query := `SELECT id, username, password, email, phone, created_at, updated_at FROM users`
	rows, err := ur.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Phone, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// CreateUser creates a new user.
func (ur *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	// SQL query to create a new user
	query := `INSERT INTO users (username, password, email, phone, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)
	return err
}

// UpdateUser updates an existing user.
func (ur *UserRepository) UpdateUser(ctx context.Context, user *model.User) error {
	// SQL query to update an existing user
	query := `UPDATE users SET username = ?, password = ?, email = ?, phone = ?, updated_at = ? WHERE id = ?`
	_, err := ur.db.ExecContext(ctx, query, user.Username, user.Password, user.Email, user.Phone, user.UpdatedAt, user.ID)
	return err
}

// DeleteUser deletes a user by ID.
func (ur *UserRepository) DeleteUser(ctx context.Context, id int) error {
	// SQL query to delete a user by ID
	query := `DELETE FROM users WHERE id = ?`
	_, err := ur.db.ExecContext(ctx, query, id)
	return err
}