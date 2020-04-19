package gohammer

import (
	"fmt"
	"os"
	"time"
)

func outResponse(response HammerResponse) {
	// fmt.Fprintf(out, "%s,%d,%d\n",
	// 	response.timestamp,
	// 	response.status,
	// 	response.latency,
	// )
	fmt.Printf("%s,%d,%d\n",
		response.Timestamp,
		response.Status,
		response.Latency,
	)
}

func outSummary(summary runSummary) {
	fmt.Println("--")
	fmt.Println("Summary")
	fmt.Println("--")
	fmt.Printf("p99: %d\n", summary.p99)
	fmt.Printf("p95: %d\n", summary.p95)
	fmt.Printf("p90: %d\n", summary.p90)
	fmt.Printf("p50: %d\n", summary.p50)
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
