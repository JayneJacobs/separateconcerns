package businesslayer

import (
	"context"
	"errors"
	"strings"
)

// Business Logic
type RegisterParams struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func (rp *RegisterParams) Validate() error {
	if rp.Email == "" {
		return errors.New("Email cannot be empty")
	}

	if !strings.ContainsRune(rp.Email, '@') {
		return errors.New("Email must include an '@' symbol")
	}

	if rp.Name == "" {
		return errors.New("Name cannot be empty")
	}

	return nil
}

type UserService interface {
	// Register may return an ErrEmailExists error
	Register(context.Context, *RegisterParams) error
	// GetByEmail may return an ErrUserNotFound error
	GetByEmail(context.Context, string) (*User, error)
}

var ErrEmailExists = errors.New("Email is already in use")

type UserServiceImpl struct {
	userStorage accesslayer.UserStorer
}

func NewUserServiceImpl(us accesslayer.UserStorer) *UserServiceImpl {
	return &UserServiceImpl{
		userStorage: us,
	}
}

func (us *UserServiceImpl) Register(ctx context.Context, params *RegisterParams) error {
	_, err := us.userStorage.Get(ctx, params.Email)
	if err == nil {
		return ErrEmailExists
	} else if err != ErrUserNotFound {
		return err
	}

	return us.userStorage.Save(ctx, &User{
		Email: params.Email,
		Name:  params.Name,
	})
}

func (us *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	return us.userStorage.Get(ctx, email)
}