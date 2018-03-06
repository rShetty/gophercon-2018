package main

import (
	"log"
	"net/http"

	_ "github.com/bcicen/grmon"
)

func main() {
	http.HandleFunc("/hello", http.HandlerFunc(hello))

	log.Print("Started server on port :8000")
	http.ListenAndServe(":8000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("Making request to slow server on port:", 9999)

	httpClient := http.Client{}
	_, err := httpClient.Get("http://localhost:9999/drivers")
	if err != nil {
		log.Print("failed with error: %s", err)
		return
	}

	w.Write([]byte("Hello Gophercon 2018"))
}
