package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/drivers", http.HandlerFunc(drivers))
	log.Print("Started slow server on port :9999")
	http.ListenAndServe(":9999", nil)
}

func drivers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(100 * time.Second)
}
