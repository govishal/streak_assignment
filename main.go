package main

import (
	"fmt"
	"net/http"
	"streak_assignment/handlers"
)

func main() {
	http.HandleFunc("/find-pairs", handlers.FindPairsHandler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}