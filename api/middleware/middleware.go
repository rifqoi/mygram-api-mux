package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/handlers"
	"github.com/rifqoi/mygram-api-mux/services"
	"go.uber.org/zap"
)

type Middleware struct {
	log     *zap.Logger
	userSvc *services.UserService
}

func NewMiddleware(log *zap.Logger, userSvc *services.UserService) Middleware {
	return Middleware{
		log,
		userSvc,
	}
}

func (m *Middleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeBefore := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Now().Sub(timeBefore)
		m.log.Info(r.Method,
			durationMillis("elapsed_time_ms", elapsed),
			zap.Duration("elapsed_time_sec", elapsed),
			zap.String("url", r.URL.String()),
		)
	})
}

func durationMillis(name string, d time.Duration) zap.Field {
	return zap.Int64(name, d.Milliseconds())
}

func (m *Middleware) CORS() func(http.Handler) http.Handler {
	cors := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "PUT", "POST", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}),
		handlers.ExposedHeaders([]string{"Link"}),
		handlers.MaxAge(300),
		handlers.AllowCredentials(),
	)
	return cors
}

func (m *Middleware) RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		next.ServeHTTP(w, r)
	})
}
