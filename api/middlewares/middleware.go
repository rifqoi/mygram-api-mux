package middlewares

import (
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"go.uber.org/zap"
)

type Middleware struct {
	log *zap.Logger
}

func NewMiddleware(log *zap.Logger) Middleware {
	return Middleware{
		log,
	}
}

func (m *Middleware) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.log.Info(r.Method,
			zap.Time("time", time.Now()),
			zap.String("url", r.URL.String()),
		)
		next.ServeHTTP(w, r)
	})
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