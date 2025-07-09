package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"-" bson:"password"` // Added password field, hidden from JSON
	NIC      string             `json:"nic" bson:"nic"`
	Address  string             `json:"address" bson:"address"`
	Birthday time.Time          `json:"birthday" bson:"birthday"`
	Gender   string             `json:"gender" bson:"gender"`
	Photo    string             `json:"photo" bson:"photo"`
	Phones   []*PhoneNumber     `json:"phones" bson:"phones,omitempty"`
}

func (u *User) Validate() bool {
	return u.Name != "" && u.Email != "" && u.NIC != "" && u.Address != "" && !u.Birthday.IsZero() && u.Gender != ""
}
