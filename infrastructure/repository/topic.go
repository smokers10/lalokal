package domain

import (
	"context"
	"lalokal/domain/topic"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type topicRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func TopicRepository(db *mongo.Database) topic.Repository {
	ctx, cancel := lib.InitializeContex()
	return &topicRepository{
		collection: *db.Collection("topic"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *topicRepository) Insert(data *topic.Topic) (failure error) {
	defer r.cancel()

	user_id, _ := primitive.ObjectIDFromHex(data.UserId)
	document := bson.M{
		"title":       data.Title,
		"description": data.Description,
		"user_id":     user_id,
	}

	if _, err := r.collection.InsertOne(r.ctx, document); err != nil {
		return err
	}

	return nil
}

func (r *topicRepository) Update(data *topic.Topic) (failure error) {
	defer r.cancel()

	_id, _ := primitive.ObjectIDFromHex(data.Id)
	document := bson.M{
		"title":       data.Title,
		"description": data.Description,
	}

	if err := r.collection.FindOneAndUpdate(r.ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *topicRepository) FindByUserId(user_id string) (result []topic.Topic) {
	defer r.cancel()

	userId, _ := primitive.ObjectIDFromHex(user_id)

	cursor, err := r.collection.Find(r.ctx, bson.M{"user_id": userId})
	if err != nil {
		return []topic.Topic{}
	}

	if err := cursor.All(r.ctx, result); err != nil {
		return []topic.Topic{}
	}

	return result
}

func (r *topicRepository) FindOneById(topic_id string) (result *topic.Topic) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	if err := r.collection.FindOne(r.ctx, bson.M{"_id": topicId}).Decode(&result); err != nil {
		return &topic.Topic{}
	}

	return result
}
