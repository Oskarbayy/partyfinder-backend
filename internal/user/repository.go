// this will contain the interfaces that we want to use
package user

import (
	"context"
	"database/sql"
)

type UserRepository interface {
	RegisterUser(ctx context.Context, u User) (User, error)
	FindByEmail(ctx context.Context, email string) (User, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (*userRepositoryImpl) RegisterUser(ctx context.Context, u User) (User, error) {
	print("test")
	return User{}, nil
}

func (*userRepositoryImpl) FindByEmail(ctx context.Context, email string) (User, error) {
	return User{}, nil
}
