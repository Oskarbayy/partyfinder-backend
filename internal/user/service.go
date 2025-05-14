package user

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{repo: r}
}

// repository
func (s *UserService) RegisterUser(ctx context.Context, Name string, Email string, Password string) (User, error) {
	hashedPw, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, fmt.Errorf("hashing password: %w", err)
	}
	Password = string(hashedPw) // overwrite with the hash

	u := User{
		ID:       uuid.New().String(),
		Name:     Name,
		Email:    Email,
		Password: Password,
	}

	// hash password, etc.
	return s.repo.RegisterUser(ctx, u)
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (User, error) {
	return s.repo.FindByEmail(ctx, email)
}
