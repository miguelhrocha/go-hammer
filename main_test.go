package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestLoadGen(t *testing.T) {

	// Test Scenario
	scenario := Scenario{}
	scenario.endpoint = "http://127.0.0.1:3000/foo"
	scenario.hammer = "HTTP"

	// Execution values
	config := RunConfig{}
	config.tps = 2
	config.duration = 10

	// Expected values
	requests := 0
	totalExpected := config.tps * config.duration

	// Prepare web server
	done := make(chan bool)
	timeout := time.After(20 * time.Second)
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
	go run(config, scenario)

	select {
	case <-done:
		fmt.Println("Test completed")
	case <-timeout:
		t.Error("Test timeout")
		fmt.Printf("Requests received %d\n", requests)
	}
}
