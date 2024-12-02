package main

import (
	"fmt"
	"net/http"
)
 
type Middleware struct {}

// TODO: is authenticated middleware
func (a *application) requireAuthentication(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authenticated := a.sessionManager.GetBool(r.Context(), "isAuthenticated")
		if !authenticated {
			clientError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized access to function"))
			return
		}

		authenticated = a.sessionManager.GetBool(r.Context(), "secondAuthDone")
		if !authenticated {
			clientError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized access to function"))
			return
		}

		id := a.sessionManager.GetInt(r.Context(), "id")
		if id == 0 {
			clientError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized access to function"))
			return
		}

		exists, err := a.udb.Exists(id)
		if err != nil {
			clientError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized access to function"))
			return
		}

		if !exists {
			clientError(w, http.StatusUnauthorized, fmt.Errorf("Unauthorized access to function"))
			return
		}
		
		next.ServeHTTP(w, r)
	})
}

// TODO: ensure json middleware
func (a *application) ensureJsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers, exists := r.Header["Content-Type"]
		if !exists {
			clientError(w, http.StatusBadRequest, fmt.Errorf("ensure json content-type"))
			return
		}

		if headers[0] != "application/json" {
			clientError(w, http.StatusBadRequest, fmt.Errorf("ensure json content-type"))
			return
		}

		next.ServeHTTP(w, r)
	})
} 
