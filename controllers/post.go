package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostController struct {
	client *mongo.Client
}

// Instantiates a new PostController using a MongoDB Client.
func NewPostController(client *mongo.Client) *PostController {
	return &PostController{client}
}

// Create a new post.
// route: /posts
// method: POST
func (pc PostController) CreatePost(w http.ResponseWriter, r *http.Request) {
	p := models.Post{}
	json.NewDecoder(r.Body).Decode(&p)

	p.Id = primitive.NewObjectID()

	_, err := InsertOne(pc.client.Database("InstagramCloneAPI").Collection("posts"), p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	pj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pj)
}

// Get details of a post by it's ID.
// route: /posts/<id>
// method: GET
func (pc PostController) GetPostById(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	id := splitPath[len(splitPath)-1]

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	p := models.Post{}
	filters := bson.M{"_id": oid}

	err = FindOne(pc.client.Database("InstagramCloneAPI").Collection("posts"), filters).Decode(&p)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	pj, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(pj)
}

// Get details of all posts of an author by ID.
// route: /posts/users/<id>
// method: GET
// query: page
func (pc PostController) GetPostsOfUser(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	id := splitPath[len(splitPath)-1]
	pq := r.URL.Query().Get("page")
	if pq == "" {
		pq = "1"
	}
	// Get current page from URL query params
	page, _ := strconv.Atoi(pq)
	var perPage int64 = 5

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	filters := bson.M{"author": oid}
	options := options.Find()
	// Set options to limit and skip for pagination
	options.SetLimit(perPage)
	options.SetSkip((int64(page) - 1) * perPage)

	p := [](models.Post){}

	err = FindAll(pc.client.Database("InstagramCloneAPI").Collection("posts"), filters, options, &p)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Get total documents and total pages for pagination
	count, _ := GetTotalCount(pc.client.Database("InstagramCloneAPI").Collection("posts"), filters)
	total_pages := math.Ceil(float64(count) / float64(perPage))

	// Include necessary information to provide pagination on client side.
	rs := bson.M{
		"result":        p,
		"page":          page,
		"total_count":   count,
		"total_pages":   total_pages,
		"has_next_page": (page != int(total_pages)),
	}

	result, err := json.Marshal(rs)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
