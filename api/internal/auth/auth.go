package auth

import (
	"context"
	"crypto/rsa"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

const GOOGLE_KEY_URL = "https://www.googleapis.com/oauth2/v3/certs"

type contextKey int

const authContextKey contextKey = 0

type Validator func() (jwt.ParseOption, error)

type remoteKeyStore struct {
	keyset *jwk.AutoRefresh
}

func (rks remoteKeyStore) option() (jwt.ParseOption, error) {
	keyset, err := rks.keyset.Fetch(context.Background(), GOOGLE_KEY_URL)
	if err != nil {
		return nil, err
	}

	return jwt.WithKeySet(keyset), nil
}

func WithGoogle() Validator {
	remoteKeystore := remoteKeys(GOOGLE_KEY_URL)
	return func() (jwt.ParseOption, error) {
		keyset, err := remoteKeystore.Fetch(context.Background(), GOOGLE_KEY_URL)
		if err != nil {
			return nil, err
		}
		return jwt.WithKeySet(keyset), nil
	}
}

func WithKey(rsaKey *rsa.PublicKey) Validator {
	key := jwk.NewRSAPublicKey()
	err := key.FromRaw(rsaKey)
	if err != nil {
		panic(err)
	}

	return func() (jwt.ParseOption, error) {
		return jwt.WithVerify(jwa.RS256, key), nil
	}
}

func NewMiddleWare(validator Validator) mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			validateOption, err := validator()
			if err != nil {
				log.Println(err)
				http.Error(w, "Unknown Server Error", http.StatusInternalServerError)
				return
			}

			token, err := jwt.ParseRequest(r, validateOption)
			if err != nil {
				log.Println(err)
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			authContext := context.WithValue(r.Context(), authContextKey, token)
			authenticatedRequest := r.WithContext(authContext)

			log.Printf("%v\n", token)

			h.ServeHTTP(w, authenticatedRequest)
		})
	}
}

func AuthToken(r *http.Request) (jwt.Token, error) {
	result := r.Context().Value(authContextKey)
	tok, ok := result.(jwt.Token)
	if !ok {
		return nil, fmt.Errorf("token not found for request")
	}

	return tok, nil

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
