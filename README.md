# go-hammer

A load generator in Go.

## Why Go? 

A load generator has the ability to generate _concurrent_ requests against a defined target, and one of the nicest features in Go is actually Concurrency –– It is easy to program and also very resource efficient. [Goroutines](https://golang.org/doc/faq#goroutines) and [Channels](https://golangbot.com/channels/) are the main characters involved in it. Goroutines are basically functions that run concurrently with other functions, not to be confused with threads, goroutines are multiplexed to a limited number of OS threads and is one of the reasons why concurrency in Go is efficient. Channels, on the other hand, are like message pipes where Goroutines can communicate to each other in a race-condition-safe manner.

## TPS

There are as many Goroutines created as TPS specified. Each Goroutine is in charge of triggering a request every second. 

## Profiling

```
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go tool pprof cpu.prof
```