package repository

import (
	"context"
	model "go-fiber-app/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneRepository struct {
	collection *mongo.Collection
}

func NewPhoneRepository(db *mongo.Database) *PhoneRepository {
	collection := db.Collection("phones")
	collection.Indexes().CreateOne(context.Background(), mongo.IndexModel{
		Keys: bson.D{{Key: "user_id", Value: 1}},
	})
	return &PhoneRepository{collection: collection}
}

func (r *PhoneRepository) CreatePhone(ctx context.Context, phone *model.PhoneNumber) error {
	result, err := r.collection.InsertOne(ctx, phone)
	if err != nil {
		return err
	}
	// Set the generated ObjectID back to the phone struct
	phone.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *PhoneRepository) FindPhoneByID(ctx context.Context, id primitive.ObjectID) (*model.PhoneNumber, error) {
	var phone model.PhoneNumber
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&phone)
	if err != nil {
		return nil, err
	}
	return &phone, nil
}

func (r *PhoneRepository) UpdatePhone(ctx context.Context, phone *model.PhoneNumber) error {
	_, err := r.collection.ReplaceOne(ctx, bson.M{"_id": phone.ID}, phone)
	return err
}

func (r *PhoneRepository) DeletePhone(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}

func (r *PhoneRepository) DeletePhonesByUserID(ctx context.Context, userID primitive.ObjectID) error {
	_, err := r.collection.DeleteMany(ctx, bson.M{"user_id": userID})
	return err
}

func (r *PhoneRepository) GetPhonesByUser(ctx context.Context, userID primitive.ObjectID) ([]model.PhoneNumber, error) {
	var phones []model.PhoneNumber
	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &phones); err != nil {
		return nil, err
	}
	return phones, nil
}
