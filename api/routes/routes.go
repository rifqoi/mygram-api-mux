package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rifqoi/mygram-api-mux/api/controller"
	"github.com/rifqoi/mygram-api-mux/api/middleware"
)

const (
	PUT    = http.MethodPut
	GET    = http.MethodGet
	POST   = http.MethodPost
	DELETE = http.MethodDelete
)

type Router struct {
	m    middleware.Middleware
	user *controller.UserController
}

func NewRouter(m middleware.Middleware, user *controller.UserController) *Router {
	return &Router{m, user}
}

func (r *Router) Run() {
	mux := mux.NewRouter()
	mux.Use(r.m.CORS())
	mux.Use(r.m.LogRequest)

	mux.HandleFunc("/users/register", r.user.RegisterUser).Methods(POST)
	mux.HandleFunc("/users/login", r.user.Login).Methods(POST)
	mux.Handle("/check", r.m.Auth(http.HandlerFunc(r.user.Check))).Methods(GET)

	log.Println("Server running at port 8000")
	http.ListenAndServe(":8000", r.m.RemoveTrailingSlash(mux))
}
