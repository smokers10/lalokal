package repository

import (
	"context"
	"lalokal/domain/user"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection mongo.Collection
	ctx        context.Context
	cancel     context.CancelFunc
}

func UserRepository(db *mongo.Database) user.Repository {
	ctx, cancel := lib.InitializeContex()
	return &userRepository{
		collection: *db.Collection("user"),
		ctx:        ctx,
		cancel:     cancel,
	}
}

func (r *userRepository) Insert(data *user.RegisterData) (inserted_id string, failure error) {
	defer r.cancel()

	document := bson.M{
		"name":         data.Name,
		"company_name": data.CompanyName,
		"email":        data.Email,
		"password":     data.Password,
	}

	inserted, err := r.collection.InsertOne(r.ctx, document)
	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *userRepository) UpdatePassword(data *user.ResetPasswordData) (failure error) {
	defer r.cancel()

	_id, _ := primitive.ObjectIDFromHex(data.UserId)
	document := bson.M{
		"password": data.Password,
	}

	if err := r.collection.FindOneAndUpdate(r.ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(data *user.User) (failure error) {
	defer r.cancel()

	_id, _ := primitive.ObjectIDFromHex(data.Id)
	document := bson.M{
		"name":         data.Name,
		"company_name": data.CompanyName,
	}

	if err := r.collection.FindOneAndUpdate(r.ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindOneByEmail(email string) (result *user.User) {
	defer r.cancel()

	if err := r.collection.FindOne(r.ctx, bson.M{"email": email}).Decode(&result); err != nil {
		return &user.User{}
	}

	return result
}

func (r *userRepository) FindOneById(user_id string) (result *user.User) {
	defer r.cancel()

	_id, _ := primitive.ObjectIDFromHex(user_id)

	if err := r.collection.FindOne(r.ctx, bson.M{"_id": _id}).Decode(&result); err != nil {
		return &user.User{}
	}

	return result
}
