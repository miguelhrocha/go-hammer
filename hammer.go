package main

import (
	"net/http"
	"sync"
	"time"
)

// Response is information about http response
type HammerResponse struct {
	latency   int // milliseconds
	status    int
	timestamp time.Time
	failed    bool
}

// Shim that decides what hammer to use
func hammer(scenario Scenario, out chan HammerResponse, wg *sync.WaitGroup) {
	if scenario.hammer == "HTTP" {
		response := httpHammer(scenario)
		out <- response
		wg.Done()
	}
}

var client *http.Client

// Hammer of type HTTP
func httpHammer(scenario Scenario) HammerResponse {
	if client == nil {
		client = new(http.Client)
		client.Timeout = time.Second * 10
	}

	// Trigger HTTP request and time it
	start := time.Now()
	res, err := client.Get(scenario.endpoint)
	end := time.Now()
	diff := end.Sub(start)

	if err != nil {
		// non-2xx response doesn't cause an error,
		// so this error means something bad happened.
		return HammerResponse{
			latency:   0,
			status:    0,
			timestamp: start.UTC(),
			failed:    true,
		}
	}

	// close response body
	defer res.Body.Close()
	return HammerResponse{
		latency:   int(diff.Milliseconds()),
		status:    res.StatusCode,
		timestamp: start.UTC(),
	}
}
