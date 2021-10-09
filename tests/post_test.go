package controllers

import (
	"net/http"
	"testing"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/controllers"
)

func TestPostController_CreatePost(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		pc   controllers.PostController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.CreatePost(tt.args.w, tt.args.r)
		})
	}
}

func TestPostController_GetPostById(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		pc   controllers.PostController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.GetPostById(tt.args.w, tt.args.r)
		})
	}
}

func TestPostController_GetPostsOfUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		pc   controllers.PostController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pc.GetPostsOfUser(tt.args.w, tt.args.r)
		})
	}
}
