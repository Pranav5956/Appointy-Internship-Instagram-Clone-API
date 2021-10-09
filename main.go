package main

import (
	"net/http"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/controllers"
)

func main() {
	// Connect to MongoDB Atlas and instantiate User and Post controllers
	client := controllers.GetClient()
	uc := controllers.NewUserController(client)
	pc := controllers.NewPostController(client)

	// Handle all routes from root using regex matching
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetRouteHandler(r.URL.Path, r.Method, uc, pc)(w, r)
	})

	// Host the server at port 5000
	http.ListenAndServe(":5000", nil)
}
