package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

// TODO: add handlers
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	var response map[string]interface{}
	json.Unmarshal([]byte(`{ "hello": "world"}`), &response)
	respondWithJSON(w, http.StatusOK, response)
}

// helper function
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
	// TODO: add routes here
	app.Router.HandleFunc("/", helloWorldHandler)
}

func (app *App) run() {
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
