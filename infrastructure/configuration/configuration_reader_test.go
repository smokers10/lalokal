package configuration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	configuration := ReadConfiguration()
	t.Log(configuration)
	t.Run("check application configuration", func(t *testing.T) {
		application := configuration.Application
		assert.NotEmpty(t, application.BaseURL)
		assert.NotEmpty(t, application.Mode)
		assert.NotEmpty(t, application.Port)
		assert.NotEmpty(t, application.Secret)
	})

	t.Run("check smtp configuration", func(t *testing.T) {
		smtp := configuration.SMTP
		assert.NotEmpty(t, smtp.Host)
		assert.NotEmpty(t, smtp.Password)
		assert.NotEmpty(t, smtp.Port)
		assert.NotEmpty(t, smtp.Sender)
		assert.NotEmpty(t, smtp.Username)
	})

	t.Run("check database configuration", func(t *testing.T) {
		database := configuration.Database
		assert.NotEmpty(t, database.DatabaseName)
		assert.NotEmpty(t, database.MaxIdleConnection)
		assert.NotEmpty(t, database.MaxPool)
		assert.NotEmpty(t, database.MinPool)
		assert.NotEmpty(t, database.MongoURI)
		assert.NotEmpty(t, database.SessionStorage)
	})
}
