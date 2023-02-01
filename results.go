package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func serveFrontendResults() {
	fs := http.FileServer(http.Dir("./assets"))
	http.Handle("/", fs)

	log.Print("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func printTotalDuration(sliceDurations []time.Duration) {
	var totalRequestsDuration time.Duration

	for _, eachTime := range sliceDurations {
		totalRequestsDuration += eachTime
	}

	fmt.Println("totalRequestsDuration", totalRequestsDuration)
}

func saveResults(filename string, requestsDuration []time.Duration) {
	var resultsInMillisecond []float32 = make([]float32, 0, len(requestsDuration))

	for _, requestDuration := range requestsDuration {
		resultsInMillisecond = append(resultsInMillisecond, float32(requestDuration)/float32(time.Millisecond))
	}

	file, err := json.MarshalIndent(resultsInMillisecond, "", " ")
	if err != nil {
		log.Fatal(err)
	}

	filepath := "assets/" + filename

	err = os.WriteFile(filepath, file, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("results saved in", filename)
}
