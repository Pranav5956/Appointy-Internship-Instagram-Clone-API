package controllers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/Pranav5956/Appointy-Internship-Instagram-Clone-API/controllers"
)

func TestMethodNotAllowedHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.MethodNotAllowedHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestPageNotFoundHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controllers.PageNotFoundHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestRouteHandlerOnMethod(t *testing.T) {
	type args struct {
		rm string
		m  string
		h  http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controllers.RouteHandlerOnMethod(tt.args.rm, tt.args.m, tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouteHandlerOnMethod() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRouteHandler(t *testing.T) {
	type args struct {
		p  string
		m  string
		uc *controllers.UserController
		pc *controllers.PostController
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		{"Correct route", args{
			"/users",
			"POST",
			controllers.NewUserController(controllers.GetClient()),
			controllers.NewPostController(controllers.GetClient()),
		}, (controllers.NewUserController(controllers.GetClient())).CreateUser},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controllers.GetRouteHandler(tt.args.p, tt.args.m, tt.args.uc, tt.args.pc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRouteHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
