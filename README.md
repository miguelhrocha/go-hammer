# go-hammer

A (work in progress) load generator in Go.

## Why Go?

A load generator has the ability to generate _concurrent_ requests against a defined target, and one of the nicest features in Go is actually Concurrency –– It is easy to program and also very resource efficient. [Goroutines](https://golang.org/doc/faq#goroutines) and [Channels](https://golangbot.com/channels/) are the main characters involved in it. Goroutines are basically functions that run concurrently with other functions; Not to be confused with threads, goroutines are actually multiplexed to a limited number of OS threads and is one of the reasons why concurrency in Go is efficient. Channels, on the other hand, are like message pipes where Goroutines can communicate to each other in a race-condition-safe manner.

## TPS

I am considering 2 approaches for generating concurrent requests with Goroutines: 

**Long living goroutines**. Create as many routines as TPS specified, where each is in charge of triggering a request every second. The challenge with this approach is that the routine waits for the request to resolve, and if the request takes longer than a second, we will not meet the desired TPS. 

**Goroutine per request**. Create a goroutine per request, where every second we spawn as many routines as TPS specified. These routines will be short-lived as they will die as soon as the request is over. Potential downside: Is it too much overhead to create many goroutines every second? 

## Reporting

TODO: Implement generation of reports that show a summary of requests but also captures raw data. Important that we show percentiles. I am thinking that each goroutine reports back its data using Channels to a metrics-dedicated routine. 

## Profiling

```
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go tool pprof cpu.prof
```