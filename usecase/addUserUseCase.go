package usecase

import (
	"context"
	"time"
)

type AddUserUseCase struct {
}

func (s *AddUserUseCase) Execute(u *User) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

}
