package entities

import (
	"time"

	"shadow-id/pkg/types"
)

// User represents a user entity in the domain
type User struct {
	ID        types.ID    `json:"id"`
	Name      string      `json:"name"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// NewUser creates a new user entity
func NewUser(name, email string) *User {
	now := time.Now()
	return &User{
		ID:        types.NewID(),
		Name:      name,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// UpdateName updates the user's name
func (u *User) UpdateName(name string) {
	u.Name = name
	u.UpdatedAt = time.Now()
}

// UpdateEmail updates the user's email
func (u *User) UpdateEmail(email string) {
	u.Email = email
	u.UpdatedAt = time.Now()
}

// Validate validates the user entity
func (u *User) Validate() error {
	if u.Name == "" {
		return ErrInvalidUserName
	}
	if u.Email == "" {
		return ErrInvalidUserEmail
	}
	return nil
}
