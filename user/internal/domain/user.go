package domain

import (
	"errors"
)

type User struct {
	Username string
	Password string
}

type UserRegister interface {
	Register(User) error
}

type UserAuthenticator interface {
	Auth(username, password string) (string, error)
}

type UserAuthorizer interface {
	Authorize(token string) (string, error)
}

type UserStore interface {
	Save(User) error
	Get(string) (User, error)
}

var (
	ErrInvalidUser       = errors.New("invalid user")
	ErrAuthFailed        = errors.New("auth failed")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrUserNotFound      = errors.New("user not found")
)

func NewUser(username, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}

func (u *User) IsValid() bool {
	return u.Username != "" && u.Password != ""
}
