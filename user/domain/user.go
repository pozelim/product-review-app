package domain

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

type UserRepository interface {
	Save(User) error
	Get(string) (User, error)
}

func NewUser(username, password string) User {
	return User{
		Username: username,
		Password: password,
	}
}

func (u *User) IsValid() bool {
	return u.Username != "" && u.Password != ""
}
