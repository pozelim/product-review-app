package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

type UserService struct {
	UserStore
	secret string
}

func NewUserService(store UserStore, secret string) *UserService {
	return &UserService{store, secret}
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

func (s *UserService) Auth(username, password string) bool {
	user, err := s.Get(username)
	if err != nil {
		return false
	}
	decryptedPassword, err := s.decryptPassword(user.Password)
	if err != nil {
		return false
	}
	return decryptedPassword == password
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
