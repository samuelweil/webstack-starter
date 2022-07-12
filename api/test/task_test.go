package e2e

import (
	"encoding/json"
	"io"
	"net/http"
	"testing"
	"weil/webstack/api/internal/task"
	tc "weil/webstack/api/test/test_client"
)

func TestGetAllTasks(t *testing.T) {

	check := func(e error) {
		if e != nil {
			t.Fatal(e)
		}
	}

	c := tc.New()

	resp, err := c.Get("/brands")
	check(err)

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expect 200 OK, got %d", resp.StatusCode)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	check(err)

	var result []task.Task

	err = json.Unmarshal(body, &result)
	check(err)

	if len(result) != 0 {
		t.Errorf("Expecting empty array, got %d elements", len(result))
	}
}
