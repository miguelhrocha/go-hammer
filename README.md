# go-hammer

A load generator in Go.

## Why Go? 

A load generator has the ability to generate concurrent requests against a defined target. One of the nicest features in Go is actually Concurrency; It's easy to program and also very resource efficient.  Concurrency in Go is done via [goroutines](https://golang.org/doc/faq#goroutines) and [Channels](https://golangbot.com/channels/). Goroutines are basically functions that run concurrently with other functions, kind of like threads, however, they are not implemented as such. Goroutines are multiplexed to a limited number of OS threads and is one of the reasons why concurrency in Go is efficient. Channels are like message buses where Goroutines can communicate to each other in a race-condition-safe manner.

## TPS

There are as many Goroutines created as TPS specified. Each Goroutine is in charge of triggering a request every second. 

## Profiling

```
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go tool pprof cpu.prof
```