package main

import (
	"flag"
	"fmt"
	"log"
	"time"
)

func main() {

	resultFlag := flag.Bool("result", false, "call this flag to up the server")
	host := flag.String("host", "", "set the server that will be receive the request")
	concurrencyFlag := flag.Int("concurrency", 1, "how many requests are in concurrency?")
	totalRequestsFlag := flag.Int("requests", 1, "how many requests do you want?")
	delayFlag := flag.Int("delay", 0, "insert a delay between the sync requests (ns)")
	providerFlag := flag.String("provider", "", "insert redis or memcached")
	actionFlag := flag.String("action", "", "insert get or set")

	flag.Parse()

	if *resultFlag {
		serveFrontendResults()
		return
	}

	if *host == "" {
		log.Fatal("you have to inform a host")
	}

	requestsDuration := run(*totalRequestsFlag, *concurrencyFlag, *delayFlag, *host, *providerFlag, *actionFlag)

	printTotalDuration(requestsDuration)
	filename := *providerFlag + "-" + *actionFlag + ".json"
	saveResults(filename, requestsDuration)
}

func run(totalRequests, timesConcurrency, delay int, host, provider, action string) []time.Duration {

	var (
		timesRequests                    = totalRequests / timesConcurrency
		requestsDuration []time.Duration = make([]time.Duration, 0, timesRequests)
	)

	fmt.Println("send", timesRequests, "requests with", timesConcurrency, "concurrency on", host)

	for i := 0; i < timesRequests; i++ {
		timeRequest := runConcurrencyRequest(timesConcurrency, host, provider, action)

		requestsDuration = append(requestsDuration, timeRequest)

		if delay > 0 {
			time.Sleep(time.Duration(delay) * time.Nanosecond)
		}
	}

	return requestsDuration
}

func runConcurrencyRequest(times int, host, provider, action string) time.Duration {
	if times == 1 {
		return cacheCall(host, provider, action)
	}

	var chanDurations = make(chan time.Duration, times)

	for i := 0; i < times; i++ {
		go func() {
			chanDurations <- cacheCall(host, provider, action)
		}()
	}

	var totalRequestsDuration time.Duration
	for i := 0; i < times; i++ {
		totalRequestsDuration += <-chanDurations
	}

	return totalRequestsDuration
}
