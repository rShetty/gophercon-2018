package main

import (
	"log"
	"net/http"
	"time"

	_ "github.com/bcicen/grmon"
	"github.com/gojektech/heimdall"
)

func main() {
	http.HandleFunc("/hello", http.HandlerFunc(hello))
	log.Print("Started server on port :8001")
	http.ListenAndServe(":8001", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("Making request to slow server on port:", 9999)

	httpClient := heimdall.NewHTTPClient(1 * time.Millisecond)
	_, err := httpClient.Get("http://localhost:9999/drivers", http.Header{})
	if err != nil {
		log.Print("failed with error: %s", err)
		return
	}

	w.Write([]byte("Hello Gophercon"))
}
