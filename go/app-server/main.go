package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Define a route for the "Hello, World!" message
	router.HandleFunc("/hello", HelloWorldHandler).Methods("GET")

	// http.Handle("/", router)

	port := os.Getenv("PORT")

	fmt.Println("Server is running on port " + port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Errorf("Error while initializing server: %+v", err)
	}
}

// HelloWorldHandler handles requests to the "/hello" route
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("In: Hello route....")

	defer fmt.Println("Out: Hello route....")

	fmt.Fprint(w, "Hello, World!")
}
