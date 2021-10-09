package controllers

import (
	"net/http"
	"testing"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/controllers"
)

func TestUserController_CreateUser(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		uc   controllers.UserController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uc.CreateUser(tt.args.w, tt.args.r)
		})
	}
}

func TestUserController_GetUserById(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		uc   controllers.UserController
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.uc.GetUserById(tt.args.w, tt.args.r)
		})
	}
}
