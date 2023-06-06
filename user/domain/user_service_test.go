package domain_test

import (
	"testing"

	"github.com/pozelim/product-review-app/common"
	"github.com/pozelim/product-review-app/user/adapters/repositories/inmemory"
	"github.com/pozelim/product-review-app/user/domain"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Register(t *testing.T) {
	type fields struct {
		userStore domain.UserRepository
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
			common.ErrInvalidResource,
		},
		{
			"Should not return error if user store returns error",
			fields{
				userStore: &MockUserStore{
					err: common.ErrResourceAlreadyExists,
				},
			},
			args{
				user: domain.User{
					Username: "test",
					Password: "test",
				},
			},
			common.ErrResourceAlreadyExists,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := domain.NewUserService(tt.fields.userStore, "6368616e676520746869732070617373", []byte("tokenSigningKey"))
			assert.Equal(t, tt.expectedErr, s.Register(tt.args.user))
		})
	}
}

func TestUserService_Auth(t *testing.T) {
	server := domain.NewUserService(inmemory.NewUserStore(), "6368616e676520746869732070617373", []byte("tokenSigningKey"))
	user := domain.User{
		Username: "test",
		Password: "test",
	}
	_ = server.Register(user)

	token, err := server.Auth(user.Username, user.Password)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestUserService_TokenValidation(t *testing.T) {
	server := domain.NewUserService(inmemory.NewUserStore(), "6368616e676520746869732070617373", []byte("tokenSigningKey"))
	user := domain.User{
		Username: "test",
		Password: "test",
	}
	_ = server.Register(user)

	token, err := server.Auth(user.Username, user.Password)

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestUserService_Authorize(t *testing.T) {
	server := domain.NewUserService(inmemory.NewUserStore(), "6368616e676520746869732070617373", []byte("tokenSigningKey"))
	user := domain.User{
		Username: "username",
		Password: "test",
	}
	_ = server.Register(user)
	token, _ := server.Auth(user.Username, user.Password)
	authUsername, err := server.Authorize(token)

	assert.Nil(t, err)
	assert.Equal(t, user.Username, authUsername)
}

func TestUserService_Authorize_Error(t *testing.T) {
	server := domain.NewUserService(inmemory.NewUserStore(), "6368616e676520746869732070617373", []byte("tokenSigningKey"))

	authUsername, err := server.Authorize("invalid token")

	assert.NotNil(t, err)
	assert.Empty(t, authUsername)
}
