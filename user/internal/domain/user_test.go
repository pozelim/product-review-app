package domain_test

import (
	"testing"

	"github.com/pozelim/product-review-app/user/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockUserStore struct {
	err  error
	user domain.User
}

func (s *MockUserStore) Save(user domain.User) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func (s *MockUserStore) Get(username string) (domain.User, error) {
	if s.err != nil {
		return domain.User{}, s.err
	}
	return s.user, nil
}

type MockUserRegister struct {
	err error
}

func (s *MockUserRegister) Register(user domain.User) error {
	if s.err != nil {
		return s.err
	}
	return nil
}

func TestUser_IsValid(t *testing.T) {
	type fields struct {
		Username string
		Password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			"User is valid",
			fields{
				Username: "test",
				Password: "test",
			},
			true,
		},
		{
			"User is invalid",
			fields{
				Username: "",
				Password: "",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := domain.NewUser(tt.fields.Username, tt.fields.Password)
			assert.Equal(t, u.IsValid(), tt.want)
		})
	}
}
