package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	Id              primitive.ObjectID `json:"id" bson:"_id"`
	Caption         string             `json:"caption" bson:"caption"`
	ImageURL        string             `json:"image_url" bson:"image_url"`
	PostedTimeStamp int                `json:"posted_timestamp" bson:"posted_timestamp"`
	Author          primitive.ObjectID `json:"author" bson:"author"`
}
