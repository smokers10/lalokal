package database

import "go.mongodb.org/mongo-driver/mongo"

type Contract interface {
	MongoDB() (database *mongo.Database, failure error)
}
