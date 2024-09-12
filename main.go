package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello class WINP2000")
}

func main() {
	http.HandleFunc("/", helloHandler) // Register the handler for the root path

	fmt.Println("Starting server on :5001")
	err := http.ListenAndServe(":5001", nil) // Start the server on port 8080
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}