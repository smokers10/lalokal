package session_store

import (
	"lalokal/infrastructure/configuration"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/mongodb"
)

func MongoSessionStore(config configuration.Database) *session.Store {
	mongo := mongodb.New(mongodb.Config{
		ConnectionURI: config.MongoURI,
		Database:      config.DatabaseName,
		Collection:    config.SessionStorage,
		Reset:         false,
	})

	return session.New(session.Config{Storage: mongo})
}
