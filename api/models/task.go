package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Status struct {
	Value     string `bson:"value" json:"value"`
	BgColor   string `bson:"bgColor" json:"bgColor"`
	TextColor string `bson:"textColor" json:"textColor"`
}

type Priority struct {
	Value     string `bson:"value" json:"value"`
	BgColor   string `bson:"bgColor" json:"bgColor"`
	TextColor string `bson:"textColor" json:"textColor"`
}

type Task struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Title         string               `bson:"title" json:"title"`
	Description   string               `bson:"description" json:"description"`
	Status        Status               `bson:"status,omitempty" json:"status"`
	Priority      Priority             `bson:"priority,omitempty" json:"priority"`
	Assignees     []primitive.ObjectID `bson:"assignees,omitempty" json:"assignees"`
	CreatedBy     primitive.ObjectID   `bson:"createdBy,omitempty" json:"createdBy"`
	ProjectID     primitive.ObjectID   `bson:"projectId,omitempty" json:"projectId"`
	DueDate       time.Time            `bson:"dueDate" json:"dueDate"`
	RemainingDays string               `bson:"remainingDays" json:"remainingDays"`
	CreatedAt     time.Time            `bson:"createdAt,omitempty" json:"createdAt"`
	UpdatedAt     time.Time            `bson:"updatedAt,omitempty" json:"updatedAt"`
}
