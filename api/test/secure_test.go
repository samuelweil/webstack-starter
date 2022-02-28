package e2e

import (
	"fmt"
	"net/http"
	"testing"
)

var SECURE_ENDPOINTS = []string{
	"secure/test",
}

func TestSecureEndpoints(t *testing.T) {
	for _, endpoint := range SECURE_ENDPOINTS {
		route := fmt.Sprintf("%s/%s", SERVER, endpoint)
		resp, err := http.Get(route)

		if err != nil {
			t.Errorf("error calling %s: %v", endpoint, err)
		}

		if resp.StatusCode != http.StatusUnauthorized {
			t.Errorf("%s return %d instead of 401", endpoint, resp.StatusCode)
		}

	}
}
