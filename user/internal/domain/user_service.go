package domain

type UserService struct {
	UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{store}
}

func (s *UserService) Register(user User) error {
	if !user.IsValid() {
		return ErrInvalidUser
	}
	return s.Save(user)
}
