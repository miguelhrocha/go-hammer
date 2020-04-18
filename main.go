package main

func run(config RunConfig, scenario Scenario) {
	stop := make(chan bool)
	done, responses := loadgen(config, scenario)
	s := aggregate(responses, stop)

	<-done         // wait until loadgen finishes
	stop <- true   // tell aggregator to stop
	summary := <-s // read summary from aggregator

	outSummary(summary)
}

func main() {

	/// Default scenario
	scenario := Scenario{}
	scenario.endpoint = "https://www.google.com"
	scenario.hammer = "HTTP"

	// Default execution values
	config := RunConfig{}
	config.tps = 3
	config.duration = 60

	run(config, scenario)
}
