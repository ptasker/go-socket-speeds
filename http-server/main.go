package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("HTTP Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
