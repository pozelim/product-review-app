package inmemory

import (
	"github.com/pozelim/product-review-app/user/internal/domain"
)

type UserStore struct {
	userMap map[string]domain.User
}

func NewUserStore() *UserStore {
	return &UserStore{
		userMap: make(map[string]domain.User),
	}
}

func (s *UserStore) Save(user domain.User) error {
	_, has := s.userMap[user.Username]
	if has {
		return domain.ErrUserAlreadyExists
	}
	s.userMap[user.Username] = user
	return nil
}

func (s *UserStore) Get(username string) (domain.User, error) {
	user, has := s.userMap[username]
	if !has {
		return domain.User{}, domain.ErrUserNotFound
	}
	return user, nil
}
