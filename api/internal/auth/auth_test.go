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

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

func makeTestRequest(v auth.KeyStore, token []byte) (*httptest.ResponseRecorder, error) {
	mw := auth.NewMiddleWare(v)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	authHeader := fmt.Sprintf("Bearer %s", string(token))
	req.Header.Add("Authorization", authHeader)
	if err != nil {
		return nil, err
	}

	testHandler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	resp := httptest.NewRecorder()
	testHandler.ServeHTTP(resp, req)
	return resp, nil
}

func TestReturn401NoToken(t *testing.T) {
	result, err := makeTestRequest(auth.WithGoogle(), []byte(""))
	if err != nil {
		t.Fatalf("%v", err)
	}

	status := result.Result().StatusCode
	if status != http.StatusUnauthorized {
		t.Fatalf("expected 401 when there's no token, got %d", status)
	}
}

func TestValidKey(t *testing.T) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	result, err := makeTestRequest(auth.WithKey(&key.PublicKey), testToken(key))

	if err != nil {
		t.Fatalf("test request failed with %v", err)
	}

	status := result.Result().StatusCode
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
