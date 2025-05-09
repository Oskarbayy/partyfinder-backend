// this will contain the interfaces that we want to use
package users

import "context"

type UserStore interface {
	Create(ctx context.Context, u *User) (*User, error)
	FindByEmail(ctx context.Context, email string) (*User, error)
}
