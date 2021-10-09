package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// This is the schema of the User document in MongoDB.
// Id is the object ID given by MongoDB.
// Name, Email and Password are all strings.
// Password will be hashed using Bcrypt in the server.
type User struct {
	Id       primitive.ObjectID `json:"id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Email    string             `json:"email" bson:"email"`
	Password string             `json:"password" bson:"password"`
}
