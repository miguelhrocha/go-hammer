package gohammer

import (
	"net/http"
	"sync"
	"time"
)

// Hammer defines functions to be implemented by hammers
type Hammer interface {
	Hit() HammerResponse
}

// HammerResponse is information about a hammer response
type HammerResponse struct {
	Latency   int // milliseconds
	Status    int
	Timestamp time.Time
	Failed    bool
}

func useHammer(h Hammer, out chan HammerResponse, wg *sync.WaitGroup) {
	response := h.Hit()
	out <- response
	wg.Done()
}

// Built-in Hammers
//
// The following are definitions and implementations of the
// built-in hammers this library offers. Anyone can develop
// their own custom hammers but for quick usage, they can use
// the built-in ones provided below.

// HTTPHammer built-in for http requests
type HTTPHammer struct {
	Endpoint string
	Method   string
}

var client *http.Client

// Hit method for HTTPHammer
func (h HTTPHammer) Hit() HammerResponse {
	if client == nil {
		client = new(http.Client)
		client.Timeout = time.Second * 10
	}

	// Trigger HTTP request and time it
	start := time.Now()
	res, err := client.Get(h.Endpoint)
	end := time.Now()
	diff := end.Sub(start)

	if err != nil {
		// non-2xx response doesn't cause an error,
		// so this error means something bad happened.
		return HammerResponse{
			Latency:   0,
			Status:    0,
			Timestamp: start.UTC(),
			Failed:    true,
		}
	}

	// close response body
	defer res.Body.Close()
	return HammerResponse{
		Latency:   int(diff.Milliseconds()),
		Status:    res.StatusCode,
		Timestamp: start.UTC(),
	}
}
