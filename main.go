package main

import (
	"fmt"
)

func aggregate(ch chan HammerResponse) {
	// createResultsDir()
	// out := createResultsFile()

	// CSV header
	// fmt.Fprintf(out, "Timestamp,Status,Latency\n")

	// Keep results in-memory
	var latencies []int64

	// Listen for messages in channel and write results
	for message := range ch {
		latencies = append(latencies, message.latency)
		// fmt.Fprintf(out, "%s,%d,%d\n",
		// 	message.timestamp,
		// 	message.status,
		// 	message.latency,
		// )
		fmt.Printf("%s,%d,%d\n",
			message.timestamp,
			message.status,
			message.latency,
		)
	}
}

func run(config RunConfig, scenario Scenario) {
	out := make(chan HammerResponse)
	done := make(chan bool)

	go loadgen(config, scenario, out, done)
	go aggregate(out)

	<-done
	fmt.Println("Good bye")
}

func main() {

	/// Default scenario
	scenario := Scenario{}
	scenario.endpoint = "https://www.google.com"
	scenario.hammer = "HTTP"

	// Default execution values
	config := RunConfig{}
	config.tps = 1
	config.duration = 10

	run(config, scenario)
}
