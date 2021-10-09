package controllers

import (
	"net/http"
	"regexp"
)

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte("method not allowed"))
}

func PageNotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("page not found"))
}

// RouteHandlerOnMethod is a helper function that returns the handler function if the
// request is of the intended method. If not, it returns MethodNotAllowedHandler.
func RouteHandlerOnMethod(rm string, m string, h http.HandlerFunc) http.HandlerFunc {
	if rm == m {
		return h
	} else {
		return MethodNotAllowedHandler
	}
}

// GetRouteHandler handles incoming request by matching it against regex strings and returns
// appropriate function handlers (http.HandlerFunc).
func GetRouteHandler(p string, m string, uc *UserController, pc *PostController) http.HandlerFunc {
	routeCreateUser := regexp.MustCompile(`^/users$`)
	routeGetUserById := regexp.MustCompile(`^/users/(\w+)$`)
	routeCreatePost := regexp.MustCompile(`^/posts$`)
	routeGetPostById := regexp.MustCompile(`^/posts/(\w+)$`)
	routeGetPostsOfUser := regexp.MustCompile(`^/posts/users/(\w+)$`)

	switch {
	case routeCreateUser.MatchString(p):
		return RouteHandlerOnMethod(m, "POST", uc.CreateUser)
	case routeGetUserById.MatchString(p):
		return RouteHandlerOnMethod(m, "GET", uc.GetUserById)
	case routeCreatePost.MatchString(p):
		return RouteHandlerOnMethod(m, "POST", pc.CreatePost)
	case routeGetPostById.MatchString(p):
		return RouteHandlerOnMethod(m, "GET", pc.GetPostById)
	case routeGetPostsOfUser.MatchString(p):
		return RouteHandlerOnMethod(m, "GET", pc.GetPostsOfUser)
	default:
		return PageNotFoundHandler
	}
}
