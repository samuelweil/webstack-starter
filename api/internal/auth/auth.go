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

type auth interface {
	verify(token string) (jwt.Token, error)
}

type remoteKeyAuth struct {
	keyset *jwk.AutoRefresh
}

func (a remoteKeyAuth) verify(token string) (jwt.Token, error) {
	keySet, err := a.keyset.Fetch(context.Background(), GOOGLE_KEY_URL)
	if err != nil {
		return nil, err
	}

	return jwt.Parse([]byte(token), jwt.WithKeySet(keySet))
}

func NewMiddleWare() mux.MiddlewareFunc {
	authService := remoteKeyAuth{
		keyset: googleKeySet(),
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr, err := getBearerToken(r.Header.Get("Authorization"))
			if err != nil {
				http.Error(w, "Invalid auth token", http.StatusUnauthorized)
				return
			}

			token, err := authService.verify(tokenStr)
			if err != nil {
				http.Error(w, "Unknown Server Error", http.StatusInternalServerError)
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
	return remoteKeys(GOOGLE_KEY_URL)
}

func remoteKeys(url string) *jwk.AutoRefresh {
	ctx := context.Background()
	keys := jwk.NewAutoRefresh(ctx)
	keys.Configure(url, jwk.WithMinRefreshInterval(15*time.Minute))

	_, err := keys.Refresh(ctx, url)
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
