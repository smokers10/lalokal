package database

import (
	"lalokal/infrastructure/lib"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMongo(t *testing.T) {
	mongo := MongoInit()

	db, err := mongo.MongoDB()

	t.Run("check for error", func(t *testing.T) {
		assert.Empty(t, err)
	})

	t.Run("check ping", func(t *testing.T) {
		ctx, cancel := lib.InitializeContex()
		defer cancel()

		err := db.Client().Ping(ctx, nil)

		assert.Empty(t, err)
	})
}
