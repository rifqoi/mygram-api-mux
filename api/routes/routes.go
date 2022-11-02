package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rifqoi/mygram-api-mux/api/middlewares"
	"github.com/rifqoi/mygram-api-mux/api/responses"
)

type Router struct {
	m middlewares.Middleware
}

func NewRouter(m middlewares.Middleware) *Router {
	return &Router{m}
}

func (r *Router) Run() {
	mux := mux.NewRouter()
	mux.Use(r.m.CORS())
	mux.Use(r.m.LogRequest)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]interface{}{
			"asd": "asd",
		}
		responses.SuccessResponse(w, data, nil)
	}).Methods(http.MethodGet)

	log.Println("Server running at port 8000")
	http.ListenAndServe(":8000", mux)
}
