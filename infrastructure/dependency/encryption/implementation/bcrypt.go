package implementation

import (
	"lalokal/infrastructure/dependency/encryption"

	"golang.org/x/crypto/bcrypt"
)

type implementation struct{}

func Bcrypt() encryption.Contract {
	return &implementation{}
}

func (b *implementation) Hash(plaintext string) (hashed_string string) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(plaintext), bcrypt.DefaultCost)
	return string(hashed)
}

func (b *implementation) Compare(hashed_text string, plain_text string) (is_correct bool) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed_text), []byte(plain_text)); err != nil {
		return false
	}

	return true
}
