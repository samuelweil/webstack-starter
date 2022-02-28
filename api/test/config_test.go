package e2e

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

func parse(resp *http.Response, into interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, into)
	if err != nil {
		return err
	}

	return nil
}

func TestConfig(t *testing.T) {
	route := fmt.Sprintf("%s/api/config", SERVER)

	resp, err := http.Get(route)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected response status 200, got %d", resp.StatusCode)
	}

	payload := make(map[string]string)
	err = parse(resp, &payload)

	if err != nil {
		t.Fatal(err)
	}

	_, ok := payload["clientId"]
	if !ok {
		t.Fatal("clientId is not provided")
	}
}
