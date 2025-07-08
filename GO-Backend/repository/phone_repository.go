package repository

import (
	"context"
	"fmt"
	model "go-fiber-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneRepository struct {
	db *mongo.Database
}

func NewPhoneRepository(db *mongo.Database) *PhoneRepository {
	return &PhoneRepository{db: db}
}

func (r *PhoneRepository) CreatePhone(ctx context.Context, phone *model.PhoneNumber) error {
	collection := r.db.Collection("phones")
	phone.ID = primitive.NewObjectID()

	_, err := collection.InsertOne(ctx, phone)
	if err != nil {
		return fmt.Errorf("error creating phone: %w", err)
	}

	return nil
}

func (r *PhoneRepository) GetPhonesByUser(ctx context.Context, userID primitive.ObjectID) ([]*model.PhoneNumber, error) {
	collection := r.db.Collection("phones")

	cursor, err := collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, fmt.Errorf("error finding phones: %w", err)
	}
	defer cursor.Close(ctx)

	var phones []*model.PhoneNumber
	for cursor.Next(ctx) {
		var phone model.PhoneNumber
		if err := cursor.Decode(&phone); err != nil {
			return nil, fmt.Errorf("error decoding phone: %w", err)
		}
		phones = append(phones, &phone)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return phones, nil
}

func (r *PhoneRepository) UpdatePhone(ctx context.Context, phone *model.PhoneNumber) error {
	collection := r.db.Collection("phones")

	filter := bson.M{"_id": phone.ID, "user_id": phone.UserID}
	update := bson.M{"$set": bson.M{
		"number": phone.Number,
		"type":   phone.Type,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("error updating phone: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("phone not found")
	}

	return nil
}

func (r *PhoneRepository) DeletePhone(ctx context.Context, userID, phoneID primitive.ObjectID) error {
	collection := r.db.Collection("phones")

	filter := bson.M{"_id": phoneID, "user_id": userID}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf("error deleting phone: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("phone not found")
	}

	return nil
}
