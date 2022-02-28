package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"weil/webstack/api/internal/auth"

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

func logMW(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

func main() {
	log.Printf("Starting API (%s)\n", VERSION)

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(aboutHandler)
	r.Use(logMW)

	r.HandleFunc("/api/config", configHandler)

	secureRouter := r.PathPrefix("/api/secure").Subrouter()
	secureRouter.Use(auth.NewMiddleWare())
	secureRouter.HandleFunc("/{endpoint}", func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		endpoint := vars["endpoint"]

		log.Printf("Logging from secure/%s\n", endpoint)

		writeJSON("Hello secure", w)
	})

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
