package main

import (
	"fmt"

	gohammer "github.com/ferdingler/go-hammer"
)

type hammer struct{}

func (h hammer) Hit() gohammer.HammerResponse {
	fmt.Println("Hitting with my custom hammer")
	return gohammer.HammerResponse{
		Latency: 1,
	}
}

func main() {

}
