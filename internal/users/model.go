package users

import "errors"

// Domain Model
type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	// more email/password rules…
	return nil
}
