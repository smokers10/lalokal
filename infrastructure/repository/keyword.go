package repository

import (
	"fmt"
	"lalokal/domain/keyword"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type keywordRepository struct {
	collection mongo.Collection
}

func KeywordRepository(db *mongo.Database) keyword.Repository {
	return &keywordRepository{
		collection: *db.Collection("keyword"),
	}
}

func (r *keywordRepository) Insert(data *keyword.Keyword) (inserted_id string, failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)
	document := bson.M{
		"keyword":  data.Keyword,
		"topic_id": topicId,
	}

	inserted, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *keywordRepository) Delete(keyword_id string) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	keywordId, _ := primitive.ObjectIDFromHex(keyword_id)

	if err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": keywordId}).Err(); err != nil {
		return err
	}

	return nil
}

func (r *keywordRepository) FindByTopicId(topic_id string) (result []keyword.Keyword) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	cursor, err := r.collection.Find(ctx, bson.M{"topic_id": topicId})
	if err != nil {
		return []keyword.Keyword{}
	}

	if err := cursor.All(ctx, &result); err != nil {
		return []keyword.Keyword{}
	}

	return result
}

func (r *keywordRepository) Cound(topic_id string) (count int) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	c, err := r.collection.CountDocuments(ctx, bson.M{"topic_id": topicId})
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(c)
}
