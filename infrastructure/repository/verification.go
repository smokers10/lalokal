package repository

import (
	"context"
	"lalokal/domain/verification"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type verificationRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func VerificationRepository(db *mongo.Database) verification.Repository {
	ctx, cancel := lib.InitializeContex()
	return &verificationRepository{
		collection: *db.Collection("verification"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *verificationRepository) Upsert(data *verification.Verification) (failure error) {
	defer r.cancel()

	document := bson.M{
		"$set": bson.M{
			"status": "not verified",
			"secret": data.Secret,
		},
	}

	if _, err := r.collection.UpdateOne(r.ctx, bson.M{"requester_email": data.RequesterEmail}, document, options.Update().SetUpsert(true)); err != nil {
		return err
	}

	return nil
}

func (r *verificationRepository) UpdateStatus(verification_id string) (failure error) {
	defer r.cancel()

	_id, _ := primitive.ObjectIDFromHex(verification_id)
	document := bson.M{"status": "verified"}

	if err := r.collection.FindOneAndUpdate(r.ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *verificationRepository) FindOneByEmail(email string) (result *verification.Verification) {
	defer r.cancel()

	if err := r.collection.FindOne(r.ctx, bson.M{"email": email}).Decode(&result); err != nil {
		return &verification.Verification{}
	}

	return result
}
