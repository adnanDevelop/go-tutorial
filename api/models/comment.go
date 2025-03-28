package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Reply struct {
	Message   string             `bson:"message" json:"message"`
	UserID    primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
}

type Comment struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Message   string             `bson:"message" json:"message"`
	ProjectID primitive.ObjectID `bson:"projectId,omitempty" json:"projectId"`
	UserID    primitive.ObjectID `bson:"userId,omitempty" json:"userId"`
	Replies   []Reply            `bson:"replies,omitempty" json:"replies"`
	CreatedAt time.Time          `bson:"createdAt,omitempty" json:"createdAt"`
}
