package main

import (
	"fmt"

	gohammer "github.com/ferdingler/go-hammer"
)

type myCustomHammer struct{}

func (h myCustomHammer) Hit() gohammer.HammerResponse {
	fmt.Println("Hitting with my custom hammer")
	return gohammer.HammerResponse{
		Latency: 1,
	}
}

func main() {

	config := gohammer.RunConfig{
		TPS:      1,
		Duration: 10,
	}

	hammer := myCustomHammer{}
	gohammer.Run(config, hammer)
}
