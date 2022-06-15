package repository

import (
	"fmt"
	"lalokal/domain/blasting_session"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blastingSessionRepository struct {
	collection mongo.Collection
}

func BlastingSessionRepository(db *mongo.Database) blasting_session.Repository {
	return &blastingSessionRepository{
		collection: *db.Collection("blasting_session"),
	}
}

func (r *blastingSessionRepository) Insert(data *blasting_session.BlastingSession) (inserted_id string, failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)
	document := bson.M{
		"title":    data.Title,
		"message":  data.Message,
		"status":   data.Status,
		"topic_id": topicId,
	}

	result, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *blastingSessionRepository) Update(data *blasting_session.BlastingSession) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(data.Id)
	document := bson.M{
		"$set": bson.M{
			"title":   data.Title,
			"message": data.Message,
		},
	}

	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *blastingSessionRepository) FindByTopicId(topic_id string) (result []blasting_session.BlastingSession) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	cursor, err := r.collection.Find(ctx, bson.M{"topic_id": topicId})
	if err != nil {
		return []blasting_session.BlastingSession{}
	}

	if err := cursor.All(ctx, result); err != nil {
		return []blasting_session.BlastingSession{}
	}

	return result
}

func (r *blastingSessionRepository) FindById(blasting_session_id string) (result *blasting_session.BlastingSession) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(blasting_session_id)

	if err := r.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result); err != nil {
		return &blasting_session.BlastingSession{}
	}

	return result
}

func (r *blastingSessionRepository) UpdateStatus(blasting_session_id string, status string) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(blasting_session_id)
	document := bson.M{
		"$set": bson.M{
			"status": status,
		},
	}

	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *blastingSessionRepository) Count(topic_id string) (count int) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	tid, _ := primitive.ObjectIDFromHex(topic_id)

	c, err := r.collection.CountDocuments(ctx, bson.M{"topic_id": tid})
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(c)
}
