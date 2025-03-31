package middlewares

import (
	"net/http"

	"github.com/gmalheiro/playground-middleware/internal/auth"
	"github.com/gmalheiro/playground-middleware/pkg/web/response"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			response.To(w).UnauthorizedErr("missing token").SendJSON()
			return
		}
		if !auth.ValidateJWT(token) {
			response.To(w).UnauthorizedErr("invalid token").SendJSON()
			return
		}
		next.ServeHTTP(w, r)
	})
}
