package main

import (
	"fmt"
	"sort"
)

type Summary struct {
	p99 int
	p95 int
	p90 int
	p50 int
}

func aggregate(responses chan HammerResponse, stop chan bool) chan Summary {
	summary := make(chan Summary)
	go func() {
		// Holds values in-memory
		var latencies []int
		for {
			select {
			// Continue reading from responses until signaled
			// to stop by the channel.
			case response := <-responses:
				latencies = append(latencies, response.latency)
				outResponse(response)
			// Signal to stop by the main routine,
			// compute summary and report it back
			case <-stop:
				fmt.Println("Aggregator finished, summarizing")
				summary <- summarize(latencies)
			}
		}
	}()
	return summary
}

func summarize(latencies []int) Summary {
	sort.Ints(latencies)
	return Summary{
		p99: percentile(latencies, 99),
		p95: percentile(latencies, 95),
		p90: percentile(latencies, 90),
		p50: percentile(latencies, 50),
	}
}

func percentile(values []int, p float32) int {
	if len(values) == 0 {
		return 0
	}

	rank := int((p / 100) * float32(len(values)+1))
	return values[rank-1]
}
