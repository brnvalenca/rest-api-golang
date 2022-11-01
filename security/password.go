package security

import "golang.org/x/crypto/bcrypt"

type IPasswordHash interface {
	GeneratePasswordHash(password string) (string, error)
	CheckPassword(hashedPassword, passwordString string) bool
}

type MyHashPassword struct{}

func NewMyHashPassword() IPasswordHash {
	return &MyHashPassword{}
}

func (*MyHashPassword) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (*MyHashPassword) CheckPassword(hashedPassword, passwordString string) bool {
	var passwordOK bool
	passwordCheck := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordString))
	if passwordCheck != nil {
		passwordOK = true
	}
	return passwordOK
}
