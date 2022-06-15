package repository

import (
	"fmt"
	"lalokal/domain/blasting_log"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type blastingLogRepository struct {
	collection mongo.Collection
}

func BlastingLogRepository(db *mongo.Database) blasting_log.Repository {
	return &blastingLogRepository{
		collection: *db.Collection("blasting_log"),
	}
}

func (r *blastingLogRepository) Insert(data *blasting_log.BlastingLogDomain) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	blastingSessionId, _ := primitive.ObjectIDFromHex(data.BlastingSessionId)
	topicId, _ := primitive.ObjectIDFromHex(data.TopicId)

	document := bson.M{
		"status":              data.Status,
		"blasting_session_id": blastingSessionId,
		"topic_id":            topicId,
	}

	if _, err := r.collection.InsertOne(ctx, document); err != nil {
		return err
	}

	return nil
}

func (r *blastingLogRepository) FindByTopicId(topic_id string) (result []blasting_log.BlastingLogDomain) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	topicId, _ := primitive.ObjectIDFromHex(topic_id)

	cursor, err := r.collection.Find(ctx, bson.M{"topic_id": topicId})
	if err != nil {
		return []blasting_log.BlastingLogDomain{}
	}

	if err := cursor.All(ctx, result); err != nil {
		return []blasting_log.BlastingLogDomain{}
	}

	return result
}

func (r *blastingLogRepository) Count(topic_id string) (count int) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	tid, _ := primitive.ObjectIDFromHex(topic_id)

	c, err := r.collection.CountDocuments(ctx, bson.M{"topic_id": tid, "status": "sent"})
	if err != nil {
		fmt.Println(err)
		return 0
	}

	return int(c)
}
