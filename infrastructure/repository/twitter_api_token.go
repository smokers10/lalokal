package repository

import (
	"lalokal/domain/twitter_api_token"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type twitterAPITokenRepository struct {
	collection mongo.Collection
}

func TwitterAPITokenRepository(db *mongo.Database) twitter_api_token.Repository {
	return &twitterAPITokenRepository{
		collection: *db.Collection("twitter_api_token"),
	}
}

func (r *twitterAPITokenRepository) Upsert(data *twitter_api_token.TwitterAPIToken) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)
	document := bson.M{
		"$set": bson.M{
			"token":  data.Token,
			"secret": data.Secret,
		},
	}

	if _, err := r.collection.UpdateOne(ctx, bson.M{"topic_id": topicId}, document, options.Update().SetUpsert(true)); err != nil {
		return err
	}

	return nil
}

func (r *twitterAPITokenRepository) FindOneByTopicId(topic_id string) (result *twitter_api_token.TwitterAPIToken) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	if err := r.collection.FindOne(ctx, bson.M{"topic_id": topicId}).Decode(&result); err != nil {
		return &twitter_api_token.TwitterAPIToken{}
	}

	return result
}
