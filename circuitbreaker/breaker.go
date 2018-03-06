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
	log.Print("Started server on port :8004")
	http.ListenAndServe(":8004", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Print("Making request to slow server on port:", 9999)

	fallbackFn := func(err error) error {
		// Handle failure
		return err
	}

	timeoutInMillis := 10
	timeout := 10 * time.Millisecond

	hystrixConfig := heimdall.NewHystrixConfig("MyCommand", heimdall.HystrixCommandConfig{
		Timeout:                timeoutInMillis,
		MaxConcurrentRequests:  100,
		ErrorPercentThreshold:  25,
		SleepWindow:            10,
		RequestVolumeThreshold: 10,
		FallbackFunc:           fallbackFn,
	})

	httpClient := heimdall.NewHystrixHTTPClient(timeout, hystrixConfig)

	_, err := httpClient.Get("http://localhost:9999/drivers", http.Header{})
	if err != nil {
		log.Print("failed with error: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Hello Gophercon"))
}
