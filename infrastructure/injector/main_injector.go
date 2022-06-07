package injector

import (
	"lalokal/infrastructure/configuration"
	"lalokal/infrastructure/database"
	"lalokal/infrastructure/encryption"
	"lalokal/infrastructure/encryption/bcrypt"
	"lalokal/infrastructure/identifier"
	"lalokal/infrastructure/identifier/google_uuid"
	"lalokal/infrastructure/jsonwebtoken"
	"lalokal/infrastructure/jsonwebtoken/jwt"
	"lalokal/infrastructure/lib"
	"lalokal/infrastructure/mailer"
	"lalokal/infrastructure/session_store"

	"github.com/gofiber/fiber/v2/middleware/session"
)

type InjectorSolvent struct {
	Repository   repositoryCompund
	Encryption   encryption.Contract
	Identifier   identifier.Contract
	JsonWebToken jsonwebtoken.Contact
	Mailer       mailer.Contract
	Session      session.Store
}

func Injector() *InjectorSolvent {
	// start database
	db, err := database.MongoInit().MongoDB()
	if err != nil {
		panic(err)
	}

	lib.CollectionBuilder(db)

	// inject database to repository resolvent
	compund := repoCompound(db)

	// call configuration
	config := configuration.ReadConfiguration()

	return &InjectorSolvent{
		Repository:   compund,
		Encryption:   bcrypt.Bcrypt(),
		JsonWebToken: jwt.JsonWebToken(),
		Mailer:       mailer.NativeSMTP(),
		Identifier:   google_uuid.GoogleUUID(),
		Session:      *session_store.MongoSessionStore(config.Database),
	}
}
