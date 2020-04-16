package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestLoadGeneration(t *testing.T) {

	// Test params
	const tps = 1
	const holdFor = 10
	const url = "http://127.0.0.1:3000/foo"

	// Expected values
	totalExpected := tps * holdFor
	requests := 0

	// Prepare web server
	done := make(chan bool)
	timeout := time.After((holdFor * 2) * time.Second)
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		requests++
		if requests == totalExpected {
			// We are done
			fmt.Printf("Total requests received %d\n", requests)
			done <- true
		}
	})

	// Start listening for requests
	go http.ListenAndServe(":3000", nil)
	time.Sleep(2 * time.Second)

	// Start load test
	go run(tps, holdFor, url)

	select {
	case <-done:
		fmt.Println("Test completed")
	case <-timeout:
		t.Error("Test timeout")
		fmt.Printf("Requests received %d\n", requests)
	}
}
