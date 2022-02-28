package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

const GOOGLE_KEY_URL = "https://www.googleapis.com/oauth2/v3/certs"

type contextKey int

const authContextKey contextKey = 0

func NewMiddleWare() mux.MiddlewareFunc {
	autoRefreshKeys := googleKeySet()

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr, err := getBearerToken(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, "Invalid auth token", http.StatusUnauthorized)
				return
			}

			keySet, err := autoRefreshKeys.Fetch(context.Background(), GOOGLE_KEY_URL)
			if err != nil {
				log.Printf("Error loading Google certs")
				http.Error(w, "Unknown Server Error", http.StatusInternalServerError)
				return
			}

			token, err := jwt.Parse([]byte(tokenStr), jwt.WithKeySet(keySet))
			if err != nil {
				http.Error(w, "Invalid auth token", http.StatusUnauthorized)
				return
			}

			authContext := context.WithValue(r.Context(), authContextKey, token)
			authenticatedRequest := r.WithContext(authContext)

			log.Printf("%v\n", token)

			h.ServeHTTP(w, authenticatedRequest)
		})
	}
}

func GetAuthToken(r *http.Request) (jwt.Token, error) {
	result := r.Context().Value(authContextKey)
	tok, ok := result.(jwt.Token)
	if !ok {
		return nil, fmt.Errorf("token not found for request")
	}

	return tok, nil

}

func googleKeySet() *jwk.AutoRefresh {
	ctx := context.Background()
	keys := jwk.NewAutoRefresh(ctx)
	keys.Configure(GOOGLE_KEY_URL, jwk.WithMinRefreshInterval(15*time.Minute))

	_, err := keys.Refresh(ctx, GOOGLE_KEY_URL)
	if err != nil {
		panic(err)
	}

	return keys
}

func getBearerToken(authHeader string) (string, error) {
	if !strings.Contains(authHeader, "Bearer") {
		return "", fmt.Errorf("ill-formed auth token")
	}

	return authHeader[8:], nil
}
