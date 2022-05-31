package repository

import (
	"context"
	"lalokal/domain/forgot_password"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forgotPasswordRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func ForgotPasswordRepository(db *mongo.Database) forgot_password.Repository {
	ctx, cancel := lib.InitializeContex()
	return &forgotPasswordRepository{
		collection: *db.Collection("forgot_password"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *forgotPasswordRepository) Insert(data *forgot_password.ForgotPassword) (failure error) {
	defer r.cancel()

	user_id, _ := primitive.ObjectIDFromHex(data.UserId)
	document := bson.M{
		"$set": bson.M{
			"token":   data.Token,
			"user_id": user_id,
			"secret":  data.Secret,
		},
	}

	if _, err := r.collection.UpdateOne(r.ctx, bson.M{"user_id": user_id}, document); err != nil {
		return err
	}

	return nil
}

func (r *forgotPasswordRepository) FindOneByToken(token string) (result *forgot_password.ForgotPassword) {
	defer r.cancel()

	if err := r.collection.FindOne(r.ctx, bson.M{"token": token}).Decode(&result); err != nil {
		return &forgot_password.ForgotPassword{}
	}

	return result
}

func (r *forgotPasswordRepository) Delete(token string) (failure error) {
	defer r.cancel()

	if err := r.collection.FindOneAndDelete(r.ctx, bson.M{"token": token}).Err(); err != nil {
		return err
	}

	return nil
}
