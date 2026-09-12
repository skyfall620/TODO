package users

import (
	"context"
	"database/sql"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (repo *UserRepository) Create(ctx context.Context, user *User) (*User, error) {
	query := `
		INSERT INTO users (name, email, password_hash)
		VALUES ($1, $2, $3)
		RETURNING id, created_at
	`
	created := &User{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}

	err := repo.database.QueryRowContext(
		ctx,
		query,
		created.Name,
		created.Email,
		created.PasswordHash,
	).Scan(&created.ID, &created.CreatedAt)

	if err != nil {
		return nil, err
	}

	return created, nil

}

func (repo *UserRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, name, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`
	user := &User{}

	err := repo.database.QueryRowContext(
		ctx,
		query,
		email,
	).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
