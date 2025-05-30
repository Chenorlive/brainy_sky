package middleware

import (
	"fmt"
	"net/http"

	"github.com/Chenorlive/brainy/utils"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract authentication information (e.g., from headers)
		token := utils.GetTokenFromRequest(r)
		if token == "" {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("missing authentication token"))
			return
		}
		fmt.Println("Token:", token)
		// Validate the token (example with a simple token check)
		if token != "valid_token" {
			utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid authentication token"))
			return
		}

		// If authentication is successful, call the next handler
		next.ServeHTTP(w, r)
	})
}
