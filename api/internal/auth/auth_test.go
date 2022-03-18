package auth_test

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"weil/webstack/api/internal/auth"

	"github.com/gorilla/mux"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func makeTestRequest(token []byte) *http.Request {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		panic(err)
	}
	authHeader := fmt.Sprintf("Bearer %s", string(token))
	req.Header.Add("Authorization", authHeader)
	return req
}

func makeTestMW() (mux.MiddlewareFunc, *rsa.PrivateKey) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return auth.NewMiddleWare(auth.WithKey(&key.PublicKey)), key
}

func TestReturn401NoToken(t *testing.T) {
	authMW, _ := makeTestMW()
	testHandler := authMW(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	request := makeTestRequest([]byte(""))

	response := httptest.NewRecorder()
	testHandler.ServeHTTP(response, request)
	status := response.Result().StatusCode
	if status != http.StatusUnauthorized {
		t.Fatalf("expected 401 when there's no token, got %d", status)
	}
}

func TestValidKey(t *testing.T) {
	authMW, privateKey := makeTestMW()
	request := makeTestRequest(testToken(privateKey))

	testHandler := authMW(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.GetAuthToken(r)
		if err != nil {
			t.Errorf("Expected valid token to be found on request")
		}
	}))

	response := httptest.NewRecorder()
	testHandler.ServeHTTP(response, request)
	status := response.Result().StatusCode
	if status != http.StatusOK {
		t.Errorf("expected response OK, got %d", status)
	}
}

func testToken(key *rsa.PrivateKey) []byte {
	issuedAt := time.Now()
	expiresAt := issuedAt.Add(time.Minute)
	raw, err := jwt.NewBuilder().IssuedAt(issuedAt).Expiration(expiresAt).Build()
	if err != nil {
		panic(err)
	}
	pvk := jwk.NewRSAPrivateKey()
	err = pvk.FromRaw(key)

	if err != nil {
		panic(err)
	}

	result, err := jwt.Sign(raw, jwa.RS256, pvk)
	if err != nil {
		panic(err)
	}

	return result
}
