package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	NIC      string             `json:"nic" bson:"nic"`
	Address  string             `json:"address" bson:"address"`
	Birthday time.Time          `json:"birthday" bson:"birthday"`
	Gender   string             `json:"gender" bson:"gender"`
	Phones   []PhoneNumber      `json:"phones" bson:"phones,omitempty"`
}

func (u *User) Validate() bool {
	return u.Name != "" && u.Email != "" && u.NIC != "" && u.Address != "" && !u.Birthday.IsZero() && u.Gender != ""
}
