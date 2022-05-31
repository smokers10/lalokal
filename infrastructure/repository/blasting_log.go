package repository

import (
	"context"
	"lalokal/domain/blasting_log"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blastingLogRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func BlastingLogRepository(db *mongo.Database) blasting_log.Repository {
	ctx, cancel := lib.InitializeContex()
	return &blastingLogRepository{
		collection: *db.Collection("blasting_log"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *blastingLogRepository) Insert(data *blasting_log.BlastingLogDomain) (failure error) {
	defer r.cancel()

	blastingSessionId, _ := primitive.ObjectIDFromHex(data.BlastingSessionId)
	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)

	document := bson.M{
		"status":              data.Status,
		"blasting_session_id": blastingSessionId,
		"topic_id":            topicId,
	}

	if _, err := r.collection.InsertOne(r.ctx, document); err != nil {
		return err
	}

	return nil
}

func (r *blastingLogRepository) FindByTopicId(topic_id string) (result []blasting_log.BlastingLogDomain) {
	defer r.cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	cursor, err := r.collection.Find(r.ctx, bson.M{"topic_id": topicId})
	if err != nil {
		return []blasting_log.BlastingLogDomain{}
	}

	if err := cursor.All(r.ctx, result); err != nil {
		return []blasting_log.BlastingLogDomain{}
	}

	return result
}
