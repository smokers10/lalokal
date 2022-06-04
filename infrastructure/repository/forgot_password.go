package repository

import (
	"lalokal/domain/forgot_password"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type forgotPasswordRepository struct {
	collection mongo.Collection
}

func ForgotPasswordRepository(db *mongo.Database) forgot_password.Repository {
	return &forgotPasswordRepository{
		collection: *db.Collection("forgot_password"),
	}
}

func (r *forgotPasswordRepository) Insert(data *forgot_password.ForgotPassword) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	user_id, _ := primitive.ObjectIDFromHex(data.UserId)
	document := bson.M{
		"$set": bson.M{
			"token":   data.Token,
			"user_id": user_id,
			"secret":  data.Secret,
		},
	}

	if _, err := r.collection.UpdateOne(ctx, bson.M{"user_id": user_id}, document); err != nil {
		return err
	}

	return nil
}

func (r *forgotPasswordRepository) FindOneByToken(token string) (result *forgot_password.ForgotPassword) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	if err := r.collection.FindOne(ctx, bson.M{"token": token}).Decode(&result); err != nil {
		return &forgot_password.ForgotPassword{}
	}

	return result
}

func (r *forgotPasswordRepository) Delete(token string) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	if err := r.collection.FindOneAndDelete(ctx, bson.M{"token": token}).Err(); err != nil {
		return err
	}

	return nil
}
