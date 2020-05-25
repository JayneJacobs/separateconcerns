package action

import (
	"context"
	"errors"
)

// Action Layer

// ErrUserNotFound error 
var ErrUserNotFound = errors.New("User not found")

// User struct
type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UserStorer interface
type UserStorer interface {
	// Get may return an ErrUserNotFound error
	Get(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

// MemoryUserStorage struct 
type MemoryUserStorage struct {
	store map[string]*User
}
// NewMemoryUserStorage struct 
func NewMemoryUserStorage() *MemoryUserStorage {
	return &MemoryUserStorage{
		store: map[string]*User{},
	}
}

// Get func
func (ms *MemoryUserStorage) Get(ctx context.Context, email string) (*User, error) {
	if u, ok := ms.store[email]; ok {
		return u, nil
	}
	return nil, ErrUserNotFound
}
// Save func ()  {
func (ms *MemoryUserStorage) Save(ctx context.Context, user *User) error {
	ms.store[user.Email] = user
	return nil
}
