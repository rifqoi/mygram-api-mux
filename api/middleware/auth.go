package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/rifqoi/mygram-api-mux/api/responses"
	"github.com/rifqoi/mygram-api-mux/helpers"
)

func (m *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			responses.UnauthorizedRequest(w, "JWT Token must be provided.")
			return
		}
		tokenString := strings.Replace(authHeader, "Bearer ", "", -1)

		claims, err := helpers.ValidateToken(tokenString)
		if err != nil {
			responses.UnauthorizedRequest(w, err.Error())
			return
		}

		// Default unmarshal dari encoding/json itu ke float64 sehingga di cast ke float64
		// https://stackoverflow.com/questions/70705673/panic-interface-conversion-interface-is-float64-not-int64
		id := int(claims["id"].(float64))
		user, err := m.userSvc.FindUserByID(r.Context(), id)
		if err != nil {
			responses.UnauthorizedRequest(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), "claims", claims)
		ctx = context.WithValue(ctx, "userInfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
