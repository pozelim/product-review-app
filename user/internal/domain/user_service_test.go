package domain_test

import (
	"testing"

	"github.com/pozelim/product-review-app/user/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Register(t *testing.T) {
	type fields struct {
		userStore domain.UserStore
	}
	type args struct {
		user domain.User
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		expectedErr error
	}{
		{
			"Should not return error if user is valid",
			fields{
				userStore: &MockUserStore{},
			},
			args{
				user: domain.User{
					Username: "test",
					Password: "test",
				},
			},
			nil,
		},
		{
			"Should not return error if user is invalid",
			fields{
				userStore: &MockUserStore{},
			},
			args{
				user: domain.User{
					Username: "test",
					Password: "",
				},
			},
			domain.ErrInvalidUser,
		},
		{
			"Should not return error if user store returns error",
			fields{
				userStore: &MockUserStore{
					err: domain.ErrUserAlreadyExists,
				},
			},
			args{
				user: domain.User{
					Username: "test",
					Password: "test",
				},
			},
			domain.ErrUserAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := domain.NewUserService(tt.fields.userStore)
			assert.Equal(t, tt.expectedErr, s.Register(tt.args.user))
		})
	}
}
