package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/controllers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

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

func main() {
	client := GetClient()
	uc := controllers.NewUserController(client)
	pc := controllers.NewPostController(client)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetRouteHandler(r.URL.Path, r.Method, uc, pc)(w, r)
	})
	http.ListenAndServe(":5000", nil)
}
