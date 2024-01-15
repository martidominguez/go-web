package middleware

import (
	"net/http"
)

// Authenticator is a middleware that authenticates the request using a token.
type Authenticator struct {
	token string
}

// NewAuthenticator creates a new Authenticator.
func NewAuthenticator(token string) *Authenticator {
	return &Authenticator{
		token: token,
	}
}

// ValidateToken checks if the person is authorized to access the resource.
func (a *Authenticator) ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// logic before
		token := r.Header.Get("Authorization")
		if token != a.token {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// call next handler
		next.ServeHTTP(w, r)
	})
}
