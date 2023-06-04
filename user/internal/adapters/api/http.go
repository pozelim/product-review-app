package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pozelim/product-review-app/user/internal/domain"
)

const BEARER_SCHEMA = "Bearer"

type HTTPServer struct {
	userRegister  domain.UserRegister
	authenticator domain.UserAuthenticator
	authorizer    domain.UserAuthorizer
}

func NewHTTPServer(
	register domain.UserRegister,
	authenticator domain.UserAuthenticator,
	authorizer domain.UserAuthorizer,
) *HTTPServer {
	return &HTTPServer{
		userRegister:  register,
		authenticator: authenticator,
		authorizer:    authorizer,
	}
}

func (s *HTTPServer) Start() error {
	r := gin.Default()
	s.router(r)
	return r.Run()
}

func (s *HTTPServer) router(r *gin.Engine) {
	r.POST("/user/register", s.register)
	r.POST("/user/authenticate", s.authenticate)
	r.GET("/user/authorize", s.authorize)
}

func (s *HTTPServer) register(c *gin.Context) {
	var newUser domain.User

	if err := c.BindJSON(&newUser); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	if err := s.userRegister.Register(newUser); err != nil {
		c.String(http.StatusConflict, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, map[string]string{})
}

func (s *HTTPServer) authenticate(c *gin.Context) {
	var user domain.User

	if err := c.BindJSON(&user); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	token, err := s.authenticator.Auth(user.Username, user.Password)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, map[string]string{
		"token": token,
	})
}

func (s *HTTPServer) authorize(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.Trim(authHeader[len(BEARER_SCHEMA):], " ")

	username, err := s.authorizer.Authorize(tokenString)
	if err != nil {
		c.String(http.StatusUnauthorized, err.Error())
		return
	}

	c.IndentedJSON(http.StatusCreated, map[string]string{
		"userName": username,
	})
}
