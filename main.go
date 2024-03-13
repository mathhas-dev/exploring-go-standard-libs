package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func helloBytes(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Success!"))
}

func helloJson(w http.ResponseWriter, req *http.Request) {
	response := map[string]string{
		"detail": "Success!",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/helloBytes", helloBytes)
	mux.HandleFunc("/helloJson", helloJson)

	// creating a http server
	http.ListenAndServe(":8090", mux)
}
