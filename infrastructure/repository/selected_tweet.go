package repository

import (
	"context"
	"lalokal/domain/selected_tweet"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type selectedTweetRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func SelectedRepository(db *mongo.Database) selected_tweet.Repository {
	ctx, cancel := lib.InitializeContex()
	return &selectedTweetRepository{
		collection: *db.Collection("selected_tweet"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

// Delete implements selected_tweet.Repository
func (r *selectedTweetRepository) Delete(selected_tweet_id string) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(selected_tweet_id)

	if err := r.collection.FindOneAndDelete(ctx, bson.M{"_id": _id}).Err(); err != nil {
		return err
	}

	return nil
}

// FindByBlastingSessionId implements selected_tweet.Repository
func (r *selectedTweetRepository) FindByBlastingSessionId(blasting_session_id string) (result []selected_tweet.SelectedTweet) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	bs_id, _ := primitive.ObjectIDFromHex(blasting_session_id)

	cursor, err := r.collection.Find(ctx, bson.M{"blasting_session_id": bs_id})
	if err != nil {
		return []selected_tweet.SelectedTweet{}
	}

	if err := cursor.All(ctx, result); err != nil {
		return []selected_tweet.SelectedTweet{}
	}

	return result
}

// Insert implements selected_tweet.Repository
func (r *selectedTweetRepository) Insert(data *selected_tweet.SelectedTweet) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	blasting_session_id, _ := primitive.ObjectIDFromHex(data.BlastingSessionId)
	document := bson.M{
		"author_id":           data.AuthorId,
		"tweet_id":            data.TweetId,
		"text":                data.Text,
		"blasting_session_id": blasting_session_id,
	}

	if _, err := r.collection.InsertOne(ctx, document); err != nil {
		return err
	}

	return nil
}
