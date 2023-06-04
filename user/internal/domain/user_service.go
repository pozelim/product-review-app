package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

const tokenExpiration = 24 * time.Hour

type UserService struct {
	UserRepository
	secret          string
	tokenSigningKey []byte
}

func NewUserService(store UserRepository, secret string, tokenSigningKey []byte) *UserService {
	return &UserService{store, secret, tokenSigningKey}
}

func (s *UserService) Register(user User) error {
	if !user.IsValid() {
		return ErrInvalidUser
	}

	var err error
	user.Password, err = s.encryptPassword(user.Password)
	if err != nil {
		return err
	}
	return s.Save(user)
}

func (s *UserService) Auth(username, password string) (string, error) {
	user, err := s.Get(username)
	if err != nil {
		return "", errors.Wrap(ErrAuthFailed, err.Error())
	}

	decryptedPassword, err := s.decryptPassword(user.Password)
	if err != nil {
		return "", errors.Wrap(ErrAuthFailed, err.Error())
	}

	if decryptedPassword != password {
		return "", errors.Wrap(ErrAuthFailed, "invalid user or password")
	}

	token, errToken := s.generateToken(user)
	if errToken != nil {
		return "", errors.Wrap(ErrAuthFailed, errToken.Error())
	}

	return token, nil
}

func (s *UserService) encryptPassword(password string) (string, error) {
	block, err := aes.NewCipher([]byte(s.secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(password)
	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], plainText)
	return encode(cipherText), nil
}

func (s *UserService) decryptPassword(text string) (string, error) {
	block, err := aes.NewCipher([]byte(s.secret))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func (s *UserService) generateToken(user User) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "user-service",
		Subject:   user.Username,
		Audience:  []string{},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpiration)),
		NotBefore: jwt.NewNumericDate(time.Now()),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.tokenSigningKey)
}

func (s *UserService) Authorize(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		return s.tokenSigningKey, nil
	})
	if err != nil {
		return "", errors.Wrap(ErrAuthFailed, err.Error())
	}

	userName, err := parsedToken.Claims.GetSubject()
	if err != nil {
		return "", errors.Wrap(ErrAuthFailed, err.Error())
	}

	return userName, nil
}
