package injector

import (
	"lalokal/infrastructure/database"
	"lalokal/infrastructure/encryption"
	"lalokal/infrastructure/encryption/bcrypt"
	"lalokal/infrastructure/identifier"
	"lalokal/infrastructure/identifier/google_uuid"
	"lalokal/infrastructure/jsonwebtoken"
	"lalokal/infrastructure/jsonwebtoken/jwt"
	"lalokal/infrastructure/lib"
	"lalokal/infrastructure/mailer"
)

type InjectorSolvent struct {
	Repository   repositoryCompund
	Encryption   encryption.Contract
	Identifier   identifier.Contract
	JsonWebToken jsonwebtoken.Contact
	Mailer       mailer.Contract
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

	return &InjectorSolvent{
		Repository:   compund,
		Encryption:   bcrypt.Bcrypt(),
		JsonWebToken: jwt.JsonWebToken(),
		Mailer:       mailer.NativeSMTP(),
		Identifier:   google_uuid.GoogleUUID(),
	}
}
