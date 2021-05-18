package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func returnError(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	errorCode := vars["Error"]
	fmt.Fprintf(w, "Error Code passed in: %s", errorCode)
}

func main() {
	r := mux.NewRouter()

	// Handlers
	r.HandleFunc("/error/{errorCode}", returnError)

	http.ListenAndServe(":80", r)

}
