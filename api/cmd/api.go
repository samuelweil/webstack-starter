package main

import (
	"log"
	"net/http"
	"os"
	"weil/webstack/api/internal/auth"
	"weil/webstack/api/internal/ctrl"
	"weil/webstack/api/internal/task"

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
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.Use(logMW)

	r.HandleFunc("/api/config", configHandler)
	r.HandleFunc("/api/about", aboutHandler)

	secureRouter := r.PathPrefix("/api").Subrouter()
	secureRouter.Use(auth.NewMiddleWare(auth.WithGoogle()))

	taskService := task.NewService()
	secureRouter.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		allTasks := taskService.AllTasks()
		ctrl.WriteJSON(w, allTasks)
	})

	http.ListenAndServe(":8080", r)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.WriteJSON(w, aboutInfo)
}

func configHandler(w http.ResponseWriter, r *http.Request) {
	ctrl.WriteJSON(w, config)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
