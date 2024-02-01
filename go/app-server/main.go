package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define a route for the "Hello, World!" message
	router.HandleFunc("/hello", HelloWorldHandler).Methods("GET")

	// Start the web server on port 8080
	http.Handle("/", router)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

// HelloWorldHandler handles requests to the "/hello" route
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}
