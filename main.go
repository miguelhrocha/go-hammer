package main

import (
	"fmt"
	"net/http"
	"time"
)

// RequestConfig represents config of http request
type RequestConfig struct {
	url     string
	holdFor int
}

func request(thread int, config RequestConfig) {
	for i := 0; i < config.holdFor; i++ {
		start := time.Now()
		res, _ := http.Get(config.url)
		end := time.Now()

		fmt.Printf("Thread=%d Status=%s Latency=%d \n", thread, res.Status, end.Sub(start))
		time.Sleep(999 * time.Millisecond)
	}
}

func main() {
	const tps = 2
	config := RequestConfig{}
	config.url = "https://google.com"
	config.holdFor = 10

	for i := 0; i < tps; i++ {
		go request(i, config)
	}

	time.Sleep(12 * time.Second)
}
