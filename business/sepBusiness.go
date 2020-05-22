package business

import (
	"context"
	"errors"
	"strings"
)

// Business Logic

// RegisterParams struct
type RegisterParams struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}


// Validate takes pointer to RegisterParams and returns error
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

// UserService is an interface requiring Register and GetByEmail methods
type UserService interface {
	// Register may return an ErrEmailExists and error
	Register(context.Context, *RegisterParams) error
	// GetByEmail may return an ErrUserNotFound error
	GetByEmail(context.Context, string) (*access.User, error)
}

// ErrEmailExists is an error message
var ErrEmailExists = errors.New("Email is already in use")

// UserServiceImpl .... s df
type UserServiceImpl struct {
	userStorage access.UserStorer
}

// NewUserServiceImpl sldf
func NewUserServiceImpl(us action.UserStorer) *UserServiceImpl {
	return &UserServiceImpl{
		userStorage: us,
	}
}

// Register ;;;
func (us *UserServiceImpl) Register(ctx context.Context, params *RegisterParams) error {
	_, err := us.userStorage.Get(ctx, params.Email)
	if err == nil {
		return ErrEmailExists
	} 
	if err != action.ErrUserNotFound {
		return err
	}

	return us.userStorage.Save(ctx, &action.User{
		Email: params.Email,
		Name:  params.Name,
	})
}

// GetByEmail get by email
func (us *UserServiceImpl) GetByEmail(ctx context.Context, email string) (*User, error) {
	return us.userStorage.Get(ctx, email)
}