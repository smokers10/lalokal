package database

import (
	"lalokal/infrastructure/configuration"
	"lalokal/infrastructure/lib"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbImplementation struct {
	databaseConfiguration configuration.Database
}

func MongoInit() *mongodbImplementation {
	configuraton := configuration.ReadConfiguration()
	return &mongodbImplementation{databaseConfiguration: configuraton.Database}
}

func (m *mongodbImplementation) MongoDB() (database *mongo.Database, failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	// set configuration
	option := options.Client().
		ApplyURI(m.databaseConfiguration.MongoURI).
		SetMaxPoolSize(uint64(m.databaseConfiguration.MaxPool)).
		SetMinPoolSize(uint64(m.databaseConfiguration.MinPool)).
		SetMaxConnIdleTime(time.Duration(m.databaseConfiguration.MaxIdleConnection))

	// set up connection
	client, err := mongo.NewClient(option)
	if err != nil {
		return nil, err
	}

	// start connection
	client.Connect(ctx)
	db := client.Database(m.databaseConfiguration.DatabaseName)

	// return database
	return db, nil
}
