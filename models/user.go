package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User has access to the system
type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"ID"`
	Username string             `bson:"Username" json:"Username"`
	Password string             `bson:"Password" json:"Password"`
	Email    string             `bson:"Email" json:"Email"`
}
