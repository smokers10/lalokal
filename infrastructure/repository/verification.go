package repository

import (
	"fmt"
	"lalokal/domain/verification"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type verificationRepository struct {
	collection mongo.Collection
}

func VerificationRepository(db *mongo.Database) verification.Repository {
	return &verificationRepository{
		collection: *db.Collection("verification"),
	}
}

func (r *verificationRepository) Upsert(data *verification.Verification) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	document := bson.M{
		"$set": bson.M{
			"status": "not verified",
			"secret": data.Secret,
		},
	}

	if _, err := r.collection.UpdateOne(ctx, bson.M{"requester_email": data.RequesterEmail}, document, options.Update().SetUpsert(true)); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *verificationRepository) UpdateStatus(verification_id string) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(verification_id)
	document := bson.M{"$set": bson.M{"status": "verified"}}

	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (r *verificationRepository) FindOneByEmail(email string) (result *verification.Verification) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	if err := r.collection.FindOne(ctx, bson.M{"requester_email": email}).Decode(&result); err != nil {
		fmt.Println(err)
		return &verification.Verification{}
	}

	return result
}
