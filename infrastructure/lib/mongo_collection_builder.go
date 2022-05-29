package lib

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CollectionBuilder(db *mongo.Database) {
	ctx, cancel := InitializeContex()
	defer cancel()

	collections := []string{
		"blasting_log",
		"blasting_session",
		"forgot_password",
		"keyword",
		"selected_tweet",
		"topic",
		"twitter_api_token",
		"user",
		"verification",
	}

	for _, c := range collections {
		db.CreateCollection(ctx, c)
		collectionIndexBuilder(db, c)
	}
}

func collectionIndexBuilder(db *mongo.Database, collection_name string) {
	ctx, cancel := InitializeContex()
	defer cancel()

	if collection_name == "blasting_session" || collection_name == "topic_id" || collection_name == "twitter_api_token" {
		index := []mongo.IndexModel{
			{Keys: bson.M{"topic_id": 1}},
		}

		db.Collection(collection_name).Indexes().CreateMany(ctx, index)
	}

	if collection_name == "forgot_password" || collection_name == "user_id" {
		index := []mongo.IndexModel{
			{Keys: bson.M{"user_id": 1}},
		}

		db.Collection(collection_name).Indexes().CreateMany(ctx, index)
	}

	if collection_name == "selected_tweet" {
		index := []mongo.IndexModel{
			{Keys: bson.M{"blasting_session_id": 1}},
		}

		db.Collection(collection_name).Indexes().CreateMany(ctx, index)
	}

	if collection_name == "blasting_log" {
		index := []mongo.IndexModel{
			{Keys: bson.M{"blasting_session_id": 1}},
			{Keys: bson.M{"topic_id": 1}},
		}

		db.Collection(collection_name).Indexes().CreateMany(ctx, index)
	}
}
