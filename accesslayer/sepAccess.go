package accesslayer

import (
	"context"
	"errors"
)

// Action Layer
var ErrUserNotFound = errors.New("User not found")

type User struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type UserStorer interface {
	// Get may return an ErrUserNotFound error
	Get(ctx context.Context, email string) (*User, error)
	Save(ctx context.Context, user *User) error
}

type MemoryUserStorage struct {
	store map[string]*User
}

func NewMemoryUserStorage() *MemoryUserStorage {
	return &MemoryUserStorage{
		store: map[string]*User{},
	}
}

func (ms *MemoryUserStorage) Get(ctx context.Context, email string) (*User, error) {
	if u, ok := ms.store[email]; ok {
		return u, nil
	}
	return nil, ErrUserNotFound
}

func (ms *MemoryUserStorage) Save(ctx context.Context, user *User) error {
	ms.store[user.Email] = user
	return nil
}