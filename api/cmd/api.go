package main

import (
	"fmt"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
)

const VERSION = "nightly"

type About struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

type Config struct {
	ClientID string `json:"clientId"`
}

var aboutInfo About = About{
	Version: VERSION,
	Name:    "web-stack-starter-api",
}

var config Config = Config{
	ClientID: os.Getenv("CLIENT_ID"),
}

func main() {
	fmt.Printf("Starting API (%s)\n", VERSION)

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(aboutHandler)
	api := r.PathPrefix("/api").Subrouter()
	api.Path("/config").HandlerFunc(configHandler)

	http.ListenAndServe(":8080", r)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(aboutInfo, w)
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(config, w)
}

func writeJSON(i interface{}, w http.ResponseWriter) error {
	resultBytes, err := json.Marshal(i)
	if err != nil {
		return err
	}

	_, err = w.Write(resultBytes)
	return err
}
