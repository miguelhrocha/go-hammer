package main

import (
	"os"
	"time"
)

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
