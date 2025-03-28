package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Project struct {
	ID          primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title       string               `bson:"title" json:"title"`
	Description string               `bson:"description" json:"description"`
	Client      primitive.ObjectID   `bson:"client,omitempty" json:"client"`
	Teams       []primitive.ObjectID `bson:"teams,omitempty" json:"teams"`
	CreatedBy   primitive.ObjectID   `bson:"createdBy,omitempty" json:"createdBy"`
	StartDate   time.Time            `bson:"startDate" json:"startDate"`
	EndDate     time.Time            `bson:"endDate" json:"endDate"`
	Comments    []primitive.ObjectID `bson:"comments,omitempty" json:"comments"`
	Tasks       []primitive.ObjectID `bson:"tasks,omitempty" json:"tasks"`
	CreatedAt   time.Time            `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt   time.Time            `bson:"updatedAt,omitempty" json:"updatedAt"`
}
