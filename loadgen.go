package gohammer

import (
	"sync"
	"time"
)

// Loadgen is in charge of generating the required load
// based on the defined TPS and execution duration.
func loadgen(config RunConfig, h Hammer) (chan bool, chan HammerResponse) {
	done := make(chan bool)
	responses := make(chan HammerResponse)
	go func() {
		// The waitGroup is used to wait for all hammers to
		// finish executing. It's basically just a counter
		// that is incremented by TPS on every interation and then
		// decremented by each goroutine that finishes.
		var wg sync.WaitGroup
		for i := 0; i < config.Duration; i++ {
			wg.Add(config.TPS)
			for j := 0; j < config.TPS; j++ {
				go useHammer(h, responses, &wg)
			}
			time.Sleep(1000 * time.Millisecond)
		}
		wg.Wait()
		done <- true
	}()
	return done, responses
}
