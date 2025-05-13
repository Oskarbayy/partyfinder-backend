package users

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	store UserStore
}

func NewService(s UserStore) *Service {
	return &Service{store: s}
}

// repository
func (s *Service) Create(ctx context.Context, u *User) (*User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("hashing password: %w", err)
	}
	u.Password = string(hashedPw) // overwrite with the hash

	// hash password, etc.
	return s.store.Create(ctx, u)
}

func (s *Service) FindByEmail(ctx context.Context, email string) (*User, error) {
	return s.store.FindByEmail(ctx, email)
}
