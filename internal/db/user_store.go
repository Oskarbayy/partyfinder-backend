package db

import (
	"context"
	"database/sql"
	"log"

	"github.com/Oskarbayy/partyfinder-backend/internal/users"
)

// Dependency Injection:
type UserStore struct {
	db *sql.DB
}

// Constructor
func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db: db}
}

// Methods Automatically includes the 'UserStore' Interface since it contains the methods required for that...
// Create
func (s *UserStore) Create(ctx context.Context, u *users.User) (*users.User, error) {
	// INSERT ... RETURNING lets Postgres send back the new row’s fields
	row := s.db.QueryRowContext(ctx, `
        INSERT INTO users (name, email, password)
        VALUES ($1, $2, $3)
        RETURNING id, name, email, password
    `, u.Name, u.Email, u.Password)

	// Scan those returned columns right back into your struct
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return nil, err
	}

	return u, nil
}

// FindByEmail
func (s *UserStore) FindByEmail(ctx context.Context, email string) (*users.User, error) {
	var u users.User

	err := s.db.
		QueryRowContext(ctx,
			"SELECT id, name, email, password FROM users WHERE email = $1",
			email,
		).
		Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		log.Fatal(err)
	}

	return &u, nil
}
