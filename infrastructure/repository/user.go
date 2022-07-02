package repository

import (
	"lalokal/domain/user"
	"lalokal/infrastructure/lib"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection mongo.Collection
}

func UserRepository(db *mongo.Database) user.Repository {
	return &userRepository{
		collection: *db.Collection("user"),
	}
}

func (r *userRepository) Insert(data *user.RegisterData) (inserted_id string, failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	document := bson.M{
		"name":         data.Name,
		"company_name": data.CompanyName,
		"email":        data.Email,
		"password":     data.Password,
	}

	inserted, err := r.collection.InsertOne(ctx, document)
	if err != nil {
		return "", err
	}

	return inserted.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *userRepository) UpdatePassword(data *user.ResetPasswordData) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(data.UserId)
	document := bson.M{
		"$set": bson.M{
			"password": data.Password,
		},
	}

	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Update(data *user.User) (failure error) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(data.Id)
	document := bson.M{
		"name":         data.Name,
		"company_name": data.CompanyName,
	}

	if err := r.collection.FindOneAndUpdate(ctx, bson.M{"_id": _id}, document).Err(); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindOneByEmail(email string) (result *user.User) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	if err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&result); err != nil {
		return &user.User{}
	}

	return result
}

func (r *userRepository) FindOneById(user_id string) (result *user.User) {
	ctx, cancel := lib.InitializeContex()
	defer cancel()

	_id, _ := primitive.ObjectIDFromHex(user_id)

	if err := r.collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&result); err != nil {
		return &user.User{}
	}

	return result
}
