package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// RequestConfig represents config of http request
type RequestConfig struct {
	url     string
	holdFor int
	thread  int
	wg      *sync.WaitGroup
}

func request(config RequestConfig) {
	for i := 0; i < config.holdFor; i++ {
		start := time.Now()
		res, _ := http.Get(config.url)
		end := time.Now()
		diff := end.Sub(start)

		fmt.Printf("Time=%s Thread=%d Status=%s Latency=%d \n",
			start.UTC(),
			config.thread,
			res.Status,
			diff.Milliseconds(),
		)

		// TODO: Is not a good idea to always wait 1 second
		// to generate the next request. We need to add the
		// time it took to do the last request.
		time.Sleep(1000 * time.Millisecond)
	}

	config.wg.Done()
}

func main() {
	const tps = 100

	var wg sync.WaitGroup
	wg.Add(tps)

	for i := 0; i < tps; i++ {
		config := RequestConfig{}
		config.url = "https://google.com"
		config.holdFor = 10
		config.thread = i
		config.wg = &wg
		go request(config)
	}

	wg.Wait()
}
