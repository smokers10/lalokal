package database

import (
	"context"
	"lalokal/infrastructure/configuration"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongodbImplementation struct {
	databaseConfiguration configuration.Database
}

func MongoInit() Contract {
	configuraton := configuration.ReadConfiguration()
	return &mongodbImplementation{databaseConfiguration: configuraton.Database}
}

func (m *mongodbImplementation) MongoDB() (database *mongo.Database, failure error) {
	// declaration
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// set configuration
	option := options.Client().
		ApplyURI(m.databaseConfiguration.MongoURI).
		SetMaxPoolSize(uint64(m.databaseConfiguration.MaxPool)).
		SetMinPoolSize(uint64(m.databaseConfiguration.MinPool)).
		SetMaxConnIdleTime(time.Duration(m.databaseConfiguration.MaxIdleConnection))

	// start connection
	client, err := mongo.Connect(ctx, option)
	if err != nil {
		return nil, err
	}

	// set connection database
	db := client.Database(m.databaseConfiguration.DatabaseName)

	// return connection
	return db, nil
}
