package config

import (
	"github.com/pozelim/product-review-app/user/internal/adapters/api"
	"github.com/pozelim/product-review-app/user/internal/adapters/repositories/inmemory"
	"github.com/pozelim/product-review-app/user/internal/domain"
)

type Application struct {
	httpServer *api.HTTPServer
}

func NewApplication() *Application {
	userService := domain.NewUserService(
		inmemory.NewUserStore(),
		"6368616e676520746869732070617373",
		[]byte("tokenSigningKey"),
	)

	return &Application{
		httpServer: api.NewHTTPServer(
			userService,
			userService,
			userService,
		),
	}
}

func (a *Application) Start() error {
	return a.httpServer.Start()
}
