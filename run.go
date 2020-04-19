package gohammer

// RunConfig holds details about the loadgen execution
type RunConfig struct {
	TPS      int
	Duration int
}

// Run runs a load test with a custom hammer
func Run(config RunConfig, h Hammer) {
	stop := make(chan bool)
	done, responses := loadgen(config, h)
	s := aggregate(responses, stop)

	<-done         // wait until loadgen finishes
	stop <- true   // tell aggregator to stop
	summary := <-s // read summary from aggregator

	outSummary(summary)
}
