package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// ResponseInfo represents response information
type ResponseInfo struct {
	latency    int64 // milliseconds
	httpStatus int
	timestamp  time.Time
}

func request(url string, ch chan ResponseInfo, wg *sync.WaitGroup) {
	start := time.Now()
	res, _ := http.Get(url)
	end := time.Now()
	diff := end.Sub(start)

	response := ResponseInfo{}
	response.latency = diff.Milliseconds()
	response.timestamp = start.UTC()
	response.httpStatus = res.StatusCode
	ch <- response
	wg.Done()
}

func capture(ch chan ResponseInfo) {
	for message := range ch {
		fmt.Printf("Time=%s Status=%d Latency=%d \n",
			message.timestamp,
			message.httpStatus,
			message.latency,
		)
	}
}

func main() {
	const tps = 10
	const holdFor = 10
	const url = "https://dev.api.awsdingler.com/v1/hello"

	ch := make(chan ResponseInfo)
	var wg sync.WaitGroup
	go capture(ch)

	for i := 0; i < holdFor; i++ {
		wg.Add(tps)
		for j := 0; j < tps; j++ {
			go request(url, ch, &wg)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	wg.Wait()
}
