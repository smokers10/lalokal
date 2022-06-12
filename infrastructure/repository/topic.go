package repository

import (
	"fmt"
	"lalokal/domain/topic"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type topicRepository struct {
	collection mongo.Collection
}

func TopicRepository(db *mongo.Database) topic.Repository {
	return &topicRepository{
		collection: *db.Collection("topic"),
	}
}

func (r *topicRepository) Insert(data *topic.Topic) (inserted_id string, failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	user_id, _ := primitive.ObjectIDFromHex(data.UserId)
	document := bson.M{
		"title":       data.Title,
		"description": data.Description,
		"user_id":     user_id,
	}

	inserted, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *topicRepository) Update(data *topic.Topic) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(data.Id)
	document := bson.M{
		"$set": bson.M{
			"title":       data.Title,
			"description": data.Description,
		},
	}

	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *topicRepository) FindByUserId(user_id string) (result []topic.Topic) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	userId, _ := primitive.ObjectIDFromHex(user_id)

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userId})
	if err != nil {
		return []topic.Topic{}
	}

	if err := cursor.All(ctx, &result); err != nil {
		return []topic.Topic{}
	}

	return result
}

func (r *topicRepository) FindOneById(topic_id string) (result *topic.Topic) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	if err := r.collection.FindOne(ctx, bson.M{"_id": topicId}).Decode(&result); err != nil {
		return &topic.Topic{}
	}

	return result
}
