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

type UserAutheticator interface {
	Auth(username, password string) bool
}

type UserStore interface {
	Save(User) error
	Get(string) (User, error)
}

var (
	ErrInvalidUser       = errors.New("invalid user")
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
