package main

import (
	"sync"
	"time"
)

// RunConfig holds details about the loadgen execution
type RunConfig struct {
	tps      int
	duration int
}

// Scenario holds details about the system under test
type Scenario struct {
	endpoint string
	hammer   string
}

// Loadgen is in charge of generating the required load
// based on the defined TPS and execution duration.
func loadgen(
	config RunConfig,
	scenario Scenario,
	out chan HammerResponse,
	done chan bool,
) {
	// The waitGroup is used to wait for all hammers to
	// finish executing. It's basically just a counter
	// that is incremented by TPS on every interation and then
	// decremented by each goroutine that finishes.
	var wg sync.WaitGroup
	for i := 0; i < config.duration; i++ {
		wg.Add(config.tps)
		for j := 0; j < config.tps; j++ {
			go hammer(scenario, out, &wg)
		}
		time.Sleep(1000 * time.Millisecond)
	}
	wg.Wait()
	done <- true
}
