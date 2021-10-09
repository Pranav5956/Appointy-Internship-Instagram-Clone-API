package controllers

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getContext() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	return ctx, cancel
}

func FindOne(coll *mongo.Collection, filters bson.M) *mongo.SingleResult {
	ctx, cancel := getContext()
	defer cancel()
	return coll.FindOne(ctx, filters)
}

func FindAll(coll *mongo.Collection, filters bson.M, options *options.FindOptions, target interface{}) error {
	ctx, cancel := getContext()
	defer cancel()
	cur, err := coll.Find(ctx, filters, options)

	if err != nil {
		return err
	}

	return cur.All(ctx, target)
}

func GetTotalCount(coll *mongo.Collection, filters bson.M) (int64, error) {
	ctx, cancel := getContext()
	defer cancel()
	return coll.CountDocuments(ctx, filters)
}

func InsertOne(coll *mongo.Collection, target interface{}) (*mongo.InsertOneResult, error) {
	ctx, cancel := getContext()
	defer cancel()
	return coll.InsertOne(ctx, target)
}
