package controllers

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Connects to MongoDB and returns a client to be used for connections.
func GetClient() *mongo.Client {
	clientOptions := options.Client().
		ApplyURI("mongodb+srv://admin:admin@cluster0.zbmwg.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	client.Ping(ctx, readpref.Primary())

	return client
}

// Creates a context with a timeout of 5 seconds.
func getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

// Helper function implementing the FindOne function in MongoDB
func FindOne(coll *mongo.Collection, filters bson.M) *mongo.SingleResult {
	ctx, cancel := getContext()
	defer cancel()
	return coll.FindOne(ctx, filters)
}

// Helper function implementing the FindAll function in MongoDB
func FindAll(coll *mongo.Collection, filters bson.M, options *options.FindOptions, target interface{}) error {
	ctx, cancel := getContext()
	defer cancel()
	cur, err := coll.Find(ctx, filters, options)

	if err != nil {
		return err
	}

	return cur.All(ctx, target)
}

// Helper function implementing the CountDocuments function in MongoDB
func GetTotalCount(coll *mongo.Collection, filters bson.M) (int64, error) {
	ctx, cancel := getContext()
	defer cancel()
	return coll.CountDocuments(ctx, filters)
}

// Helper function implementing the InsertOne function in MongoDB
func InsertOne(coll *mongo.Collection, target interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := getContext()
	defer cancel()
	return coll.InsertOne(ctx, target)
}
