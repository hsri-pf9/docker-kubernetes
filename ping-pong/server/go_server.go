package main

import (
	"fmt"
	"net/http"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/ping" && r.Method == "GET" {
		fmt.Fprint(w, "pong from go")
	} else {
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("Server running on port 3001...")
	http.ListenAndServe(":3001", nil)
}
