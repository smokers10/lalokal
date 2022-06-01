package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonWebToken(t *testing.T) {
	payload := map[string]interface{}{
		"name": "john doe",
		"age":  24,
	}
	jwt := JsonWebToken()
	token, err := jwt.Sign(payload)

	t.Run("check sign token", func(t *testing.T) {
		assert.Empty(t, err)
		assert.NotEmpty(t, token)
	})

	t.Run("check parsing", func(t *testing.T) {
		p, err := jwt.ParseToken(token)

		assert.Empty(t, err)
		assert.NotEmpty(t, p["name"].(string))
		assert.NotEmpty(t, p["age"].(float64))
		assert.Equal(t, "john doe", p["name"].(string))
		assert.Equal(t, 24, int(p["age"].(float64)))
	})
}
