package middlewares

import "net/http"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("x-auth-token")
		if token == "" {
			http.Error(rw, "No token Provided", http.StatusUnauthorized)
			return
		}
		// auth here
	}
}

func RequireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// check token health
		// authorize
		next(rw, r)
	}
}

func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		next(rw, r)
	}
}
