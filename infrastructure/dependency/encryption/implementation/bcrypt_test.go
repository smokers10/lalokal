package implementation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBcrypt(t *testing.T) {
	plaintext := "abcd123"
	hashed := Bcrypt().Hash(plaintext)

	t.Run("hashed should not empty", func(t *testing.T) {
		assert.NotEmpty(t, hashed)
	})

	t.Run("correct comparation", func(t *testing.T) {
		is_correct := Bcrypt().Compare(hashed, "abcd123")

		assert.Equal(t, true, is_correct)
	})

	t.Run("incorrect comparation", func(t *testing.T) {
		is_correct := Bcrypt().Compare(hashed, "efgh456")

		assert.Equal(t, false, is_correct)
	})
}
