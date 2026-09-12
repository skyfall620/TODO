package auth

import (
	"context"
	"todo/internal/users"
)

type IUsersRepository interface {
	Create(ctx context.Context, user *users.User) (*users.User, error)
	GetByEmail(ctx context.Context, email string) (*users.User, error)
}
