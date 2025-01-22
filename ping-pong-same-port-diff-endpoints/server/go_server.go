package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/go-ping" && r.Method == "GET" {
		fmt.Fprint(w, "pong from go")
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/go-ping", pingHandler)
	fmt.Println("Server running on port 8082...")
	http.ListenAndServe(":8082", nil)
}
