package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/librato"
	"log"
	"os"
	"time"
)

func startReporters() {
	libratoEmail := os.Getenv("LIBRATO_EMAIL")
	libratoToken := os.Getenv("LIBRATO_TOKEN")

	if len(libratoEmail) > 0 && len(libratoToken) > 0 {
		go func() {
			timer := 1 * time.Minute
			reporter := librato.NewReporter(metrics.DefaultRegistry, timer, libratoEmail, libratoToken, "collector", make([]float64, 0), timer)
			reporter.Run()
		}()
	} else {
		fmt.Println("Ignoring librato reporter because no email nor token")
	}
	go metrics.Log(metrics.DefaultRegistry, 60e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))
}
