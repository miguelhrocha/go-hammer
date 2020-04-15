package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Response is information about http response
type Response struct {
	latency    int64 // milliseconds
	httpStatus int
	timestamp  time.Time
}

func request(url string, ch chan Response, wg *sync.WaitGroup, client *http.Client) {
	start := time.Now()
	res, err := client.Get(url)
	end := time.Now()
	diff := end.Sub(start)

	if err != nil {
		// non-2xx response doesn't cause an error.
		panic(err)
	}

	// close response body
	defer res.Body.Close()

	// Communicate back the response details
	response := Response{}
	response.latency = diff.Milliseconds()
	response.timestamp = start.UTC()
	response.httpStatus = res.StatusCode
	ch <- response

	// Notify completion to wait group
	wg.Done()
}

func capture(ch chan Response) {
	for message := range ch {
		fmt.Printf("Time=%s Status=%d Latency=%d \n",
			message.timestamp,
			message.httpStatus,
			message.latency,
		)
	}
}

func main() {
	const tps = 240
	const holdFor = 10
	// const url = "https://dev.api.awsdingler.com/v1/hello"
	const url = "https://www.google.com"
	total := 0

	ch := make(chan Response)
	var wg sync.WaitGroup
	go capture(ch)

	client := http.Client{
		Timeout: time.Second * 10,
	}

	for i := 0; i < holdFor; i++ {
		wg.Add(tps)
		for j := 0; j < tps; j++ {
			total = total + 1
			go request(url, ch, &wg, &client)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	wg.Wait()
	fmt.Printf("Total requests=%d \n", total)
}
