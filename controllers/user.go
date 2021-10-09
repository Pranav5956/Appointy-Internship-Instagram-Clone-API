package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	client *mongo.Client
}

// Instantiates a new UserController with a MongoDB Client
func NewUserController(client *mongo.Client) *UserController {
	return &UserController{client}
}

// Create a new user.
// route: /users
// method: POST
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	u.Id = primitive.NewObjectID()
	// Hash the password using bcrypt
	hpwd, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	u.Password = string(hpwd)

	_, err = InsertOne(uc.client.Database("InstagramCloneAPI").Collection("users"), u)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}

// Get details of a user by their ID.
// route: /users/<id>
// method: GET
func (uc UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	splitPath := strings.Split(r.URL.Path, "/")
	id := splitPath[len(splitPath)-1]

	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	u := models.User{}
	filters := bson.M{"_id": oid}

	err = FindOne(uc.client.Database("InstagramCloneAPI").Collection("users"), filters).Decode(&u)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(uj)
}
