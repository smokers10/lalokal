package domain

import (
	"context"
	"lalokal/domain/blasting_session"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blastingSessionRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func BlastingSessionRepository(db *mongo.Database) blasting_session.Repository {
	ctx, cancel := lib.InitializeContex()
	return &blastingSessionRepository{
		collection: *db.Collection("blasting_session"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *blastingSessionRepository) Insert(data *blasting_session.BlastingSession) (failure error) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)
	document := bson.M{
		"title":    data.Title,
		"message":  data.Message,
		"status":   data.Status,
		"topic_id": topicId,
	}

	if _, err := r.collection.InsertOne(r.ctx, document); err != nil {
		return err
	}

	return nil
}

func (r *blastingSessionRepository) Update(data *blasting_session.BlastingSession) (failure error) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)
	_id, _ := primitive.ObjectIDFromHex(data.Id)
	document := bson.M{
		"$set": bson.M{
			"title":    data.Title,
			"message":  data.Message,
			"status":   data.Status,
			"topic_id": topicId,
		},
	}

	if err := r.collection.FindOneAndUpdate(r.ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *blastingSessionRepository) FindByTopicId(topic_id string) (result []blasting_session.BlastingSession) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	cursor, err := r.collection.Find(r.ctx, bson.M{"topic_id": topicId})
	if err != nil {
		return []blasting_session.BlastingSession{}
	}

	if err := cursor.All(r.ctx, result); err != nil {
		return []blasting_session.BlastingSession{}
	}

	return result
}

func (r *blastingSessionRepository) FindById(blasting_session_id string) (result *blasting_session.BlastingSession) {
	defer r.cancel()

	_id, _ := primitive.ObjectIDFromHex(blasting_session_id)

	if err := r.collection.FindOne(r.ctx, bson.M{"_id": _id}).Decode(&result); err != nil {
		return &blasting_session.BlastingSession{}
	}

	return result
}
