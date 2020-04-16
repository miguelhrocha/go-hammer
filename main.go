package main

import (
	"fmt"
	"net/http"
	"os"
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

func createResultsDir() {
	const dir = "results"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if os.Mkdir(dir, 0755) != nil {
			panic("failed to create results directory")
		}
	}
}

func createResultsFile() *os.File {
	os.Chdir("results")
	name := time.Now().Format(time.UnixDate)
	f, err := os.Create(name + ".csv")
	if err != nil {
		panic("failed to create results file")
	}
	return f
}

func capture(ch chan Response) {
	createResultsDir()
	out := createResultsFile()

	// CSV header
	fmt.Fprintf(out, "Timestamp,Status,Latency\n")

	// Listen for messages in channel and write results
	for message := range ch {
		fmt.Fprintf(out, "%s,%d,%d\n",
			message.timestamp,
			message.httpStatus,
			message.latency,
		)
	}
}

func run(tps int, holdFor int, url string) {
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

func main() {
	const tps = 240
	const holdFor = 10
	const url = "https://www.google.com"
	run(tps, holdFor, url)
}
