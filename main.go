package main

import (
	"fmt"
	"net/http"
	"time"
)

// Handler will handle the http request, receives name
// and then returns a json saying "hello, {name}" 
// and timestamp in unix format
func handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now().Unix()
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "World"
	}
	response := fmt.Sprintf(`{"message": "hello, %s","time":"%d"}`, name, now)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, response)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("server running on port :8090")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Printf("failed to run server: %v\n", err)
		return
	}
}
