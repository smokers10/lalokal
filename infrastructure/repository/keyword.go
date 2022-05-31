package repository

import (
	"context"
	"lalokal/domain/keyword"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type keywordRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func KeywordRepository(db *mongo.Database) keyword.Repository {
	ctx, cancel := lib.InitializeContex()
	return &keywordRepository{
		collection: *db.Collection("keyword"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *keywordRepository) Insert(data *keyword.Keyword) (inserted_id string, failure error) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)
	document := bson.M{
		"keyword":  data.Keyword,
		"topic_id": topicId,
	}

	inserted, err := r.collection.InsertOne(r.ctx, document)
	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *keywordRepository) Delete(keyword_id string) (failure error) {
	defer r.cancel()

	keywordId, _ := primitive.ObjectIDFromHex(keyword_id)

	if err := r.collection.FindOneAndDelete(r.ctx, bson.M{"_id": keywordId}).Err(); err != nil {
		return err
	}

	return nil
}

func (r *keywordRepository) FindByTopicId(topic_id string) (result []keyword.Keyword) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	cursor, err := r.collection.Find(r.ctx, bson.M{"topic_id": topicId})
	if err != nil {
		return []keyword.Keyword{}
	}

	if err := cursor.All(r.ctx, result); err != nil {
		return []keyword.Keyword{}
	}

	return result
}
