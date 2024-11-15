package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/your-project/grpc-redis-postgres/backend/services"
)

// JWTAuth middleware for JWT token authentication
func JWTAuth(next http.Handler, tokenService services.TokenService) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
			return
		}

		// Split the header value into "Bearer " and the token
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "Invalid authorization header format", http.StatusUnauthorized)
			return
		}

		// Verify the JWT token
		tokenString := tokenParts[1]
		userId, err := tokenService.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Store the user ID in the request context
		ctx := context.WithValue(r.Context(), "userId", userId)
		r = r.WithContext(ctx)

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}
