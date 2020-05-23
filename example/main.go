package main

import (
	"fmt"
	"github.com/chaseisabelle/histo"
	"time"
)

func main() {
	his, err := histo.New("histo_test_latency_seconds", "bla bla bla", []string{
		"foo",
	}, nil)

	if err != nil {
		panic(err)
	}

	//////////////////////
	// using the stopwatch
	//////////////////////
	his.Start()

	// do something

	his.Stop()
	his.Record("bar")

	println(fmt.Sprintf("took this many seconds %f", his.Duration()))

	//////////////////////////
	// using your own duration
	//////////////////////////

	start := time.Now()

	// do something

	dur := time.Since(start).Seconds()

	his.Observe(dur, "bar")

	println(fmt.Sprintf("took this many seconds %f", his.Duration()))
}
