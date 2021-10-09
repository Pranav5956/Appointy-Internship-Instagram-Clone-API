package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This is the schema of the Post document in MongoDB.
// Id is the object id given by MongoDB.
// Caption and Image URL are strings.
// PostedTimestamp is an int, which contains the unix timestamp.
// Author is the object ID pointing to a User document.
type Post struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	Caption         string             `json:"caption" bson:"caption"`
	ImageURL        string             `json:"image_url" bson:"image_url"`
	PostedTimeStamp int                `json:"posted_timestamp" bson:"posted_timestamp"`
	Author          primitive.ObjectID `json:"author" bson:"author"`
}
