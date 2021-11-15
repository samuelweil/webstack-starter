package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

const VERSION = "nightly"

type About struct {
	Version string `json:"version"`
}

var aboutInfo About = About{
	Version: "nightly",
}

func main() {
	fmt.Printf("Starting API (%s)\n", VERSION)

	r := mux.NewRouter()
	r.PathPrefix("/api").HandlerFunc(aboutHandler)

	http.ListenAndServe(":8080", r)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(aboutInfo, w)
}

func writeJSON(i interface{}, w http.ResponseWriter) error {
	resultBytes, err := json.Marshal(i)
	if err != nil {
		return err
	}

	_, err = w.Write(resultBytes)
	return err
}
