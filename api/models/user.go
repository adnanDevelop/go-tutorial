package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role string

const (
	Member         Role = "member"
	ProjectManager Role = "projectManager"
	Admin          Role = "admin"
)

type User struct {
	ID             string               `json:"id" bson:"_id,omitempty"`
	Name           string               `json:"name" bson:"name"`
	Email          string               `json:"email" bson:"email"`
	Password       string               `json:"password" bson:"password"`
	ProfilePicture string               `json:"profilePicture" bson:"profilePicture"`
	Description    string               `json:"description" bson:"description"`
	Role           Role                 `json:"role" bson:"role"`
	Deal           string               `json:"deal" bson:"deal"`
	Country        string               `json:"country" bson:"country"`
	Status         string               `json:"status" bson:"status"`
	Projects       []primitive.ObjectID `json:"projects" bson:"projects"`
	Tasks          []primitive.ObjectID `json:"tasks" bson:"tasks"`
}
