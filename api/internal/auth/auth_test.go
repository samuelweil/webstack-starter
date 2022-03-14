package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"weil/webstack/api/internal/auth"
)

func makeTestRequest(v auth.TokenVerifier, token string) (*httptest.ResponseRecorder, error) {
	mw := auth.NewMiddleWare(v)
	req, err := http.NewRequest(http.MethodGet, "/", nil)
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
	result, err := makeTestRequest(auth.WithGoogle(), "")
	if err != nil {
		t.Fatalf("%v", err)
	}

	status := result.Result().StatusCode
	if status != http.StatusUnauthorized {
		t.Fatalf("expected 401 when there's no token, got %d", status)
	}
}
