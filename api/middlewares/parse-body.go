package middlewares

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/kompere/kompere-api/cmd"
)

func ParseBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Parse JSON request body into a map
		var body map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Create a new context with the parsed request body
		ctx := context.WithValue(r.Context(), cmd.Body, body)

		// Call the next handler with the updated context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
