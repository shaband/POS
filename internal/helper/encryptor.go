package helper

import "golang.org/x/crypto/bcrypt"

type Encryptor struct{}

func RegisterEncryptor() *Encryptor {
	return &Encryptor{}
}

func (e *Encryptor) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
