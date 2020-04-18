package main

import (
	"fmt"

	gohammer "github.com/ferdingler/go-hammer"
)

func main() {
	fmt.Println("Starting load test")
	config := gohammer.RunConfig{
		TPS:      10,
		Duration: 60,
	}

	hammer := gohammer.HTTPHammer{}
	hammer.Endpoint = "https://www.google.com"
	hammer.Method = "GET"

	gohammer.Run(config, hammer)
}
