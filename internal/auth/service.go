package auth

import (
	"context"
	"database/sql"
	"errors"
	"todo/internal/users"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UsersRepository IUsersRepository
}

func NewAuthService(usersRepository IUsersRepository) *AuthService {
	return &AuthService{UsersRepository: usersRepository}
}

func (a *AuthService) Register(ctx context.Context, name, email, password string) (int64, error) {
	_, err := a.UsersRepository.GetByEmail(ctx, email)
	if err == nil {
		return -1, ErrUserExists
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return -1, err
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return -1, err
	}

	user := &users.User{
		Name:         name,
		Email:        email,
		PasswordHash: string(passwordHash),
	}

	createdUser, err := a.UsersRepository.Create(ctx, user)
	if err != nil {
		return -1, err
	}

	return createdUser.ID, nil
}

func (a *AuthService) Login(ctx context.Context, email, password string) (int64, error) {
	user, err := a.UsersRepository.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return -1, ErrUserNotFound
		}
		return -1, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return -1, ErrWrongCredetials
	}

	return user.ID, nil
}
