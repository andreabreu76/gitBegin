package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Hello World"}`))
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	log.Printf("Server started at port 8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}

}
