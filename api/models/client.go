package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Client struct {
	ID            primitive.ObjectID   `bson:"_id,omitempty" json:"id"`
	Name          string               `bson:"name" json:"name"`
	Email         string               `bson:"email" json:"email"`
	Password      string               `bson:"password,omitempty" json:"-"`
	Description   string               `bson:"description" json:"description"`
	Deal          string               `bson:"deal" json:"deal"`
	Country       string               `bson:"country" json:"country"`
	Status        string               `bson:"status" json:"status"`
	Projects      []primitive.ObjectID `bson:"projects,omitempty" json:"projects"`
	Tasks         []primitive.ObjectID `bson:"tasks,omitempty" json:"tasks"`
	CreatedBy     primitive.ObjectID   `bson:"createdBy,omitempty" json:"createdBy"`
	ClientPicture string               `bson:"clientPicture,omitempty" json:"clientPicture"`
	Active        bool                 `bson:"active" json:"active"`
	Role          string               `bson:"role" json:"role"`
}
