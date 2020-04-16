![Build](https://github.com/ferdingler/go-hammer/workflows/Build/badge.svg)

# go-hammer

A (work in progress) load generator in Go.

## Why Go?

A load generator has the ability to generate _concurrent_ requests against a defined target, and one of the nicest features in Go is actually [Concurrency](https://www.youtube.com/watch?v=cN_DpYBzKso) –– It is easy to program and also very resource efficient. [Goroutines](https://golang.org/doc/faq#goroutines) and [Channels](https://golangbot.com/channels/) are the main characters involved in it. Goroutines are basically functions that run concurrently with other functions; Not to be confused with threads, goroutines are actually multiplexed to a limited number of OS threads and is one of the reasons why concurrency in Go is efficient. Channels are message pipes where Goroutines can communicate to each other in a race-condition-safe manner.

## TPS Generation

I am considering two approaches for generating concurrent requests with Goroutines: 

**Long living goroutines**. Create as many routines as TPS specified, where each is in charge of generating a request every second. The challenge with this approach is that the goroutine waits for the request to resolve and if it takes longer than a second, it will not meet the desired TPS.

**Goroutine per request**. Create a goroutine per request where every second, there will be as many goroutines as TPS specified spawned. These routines will be short-lived as they will die as soon as the request resolves. Potential downside: Will it be too much overhead to create hundreds/thousands of goroutines every second? What if the requests are timing out and the _running active_ goroutines start to pile up?

## Reporting

The approach for reporting results would be for each goroutine to send results of the requests to a shared Channel. There is a separate goroutine in charge of listening for these results and keeping a running count and do a summary at the end. Challenge? How does this goroutine know that all requests have finished and when should it start to calculate the final results?

## Profiling

```
go test -cpuprofile cpu.prof -memprofile mem.prof -bench .
go tool pprof cpu.prof
```

## Limits

The max TPS seems to be constrained by the host OS max open files limit. The package net/http opens a socket for every HTTP request in-flight. In my macbook this seems pretty low as it breaks at ~225 TPS. Command to find this limit:

```
ulimit -n
```

Need to test on an EC2 and within a Fargate container and compare.

## Other thoughts

**Distribute requests evenly within the timeframe of a second?**. The current implementation triggers all requests in the beginning of a second, obviously each consecutive request is triggered slightly (a few microseconds?) after the previous one, but this means that all the remaining time before the second is over is empty and quiet with no requests. I suspect that distributing all of these requests evenly within the timeframe of a second would mimic real-world  traffic more realistically and would create less pressure on the system under test. Although maybe the inherent variability in the latency of the internet and networks already accomplishes this? Not sure.