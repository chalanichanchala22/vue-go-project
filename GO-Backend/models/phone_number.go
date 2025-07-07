package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhoneNumber struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Number string             `json:"number" bson:"number"`
	Type   string             `json:"type" bson:"type"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}

func (p *PhoneNumber) Validate() bool {
	return p.Number != "" && p.Type != "" && !p.UserID.IsZero()
}
