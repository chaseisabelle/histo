package main

import (
	"github.com/chaseisabelle/histo"
	"time"
)

func main() {
	histo, err := histo.New("histo_test_latency_seconds", "bla bla bla", []string{
		"foo",
	}, nil)

	if err != nil {
		panic(err)
	}

	start := time.Now()

	// do something

	dur := time.Since(start).Seconds()

	histo.Observe(dur, "bar")
}
