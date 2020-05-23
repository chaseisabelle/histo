package histo

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type Histo struct {
	vector   *prometheus.HistogramVec
	start    time.Time
	duration time.Duration
}

func New(name string, help string, labels []string, buckets []float64) (*Histo, error) {
	opts := prometheus.HistogramOpts{
		Name: name,
		Help: help,
	}

	if buckets != nil && len(buckets) > 0 {
		opts.Buckets = buckets
	}

	vector := prometheus.NewHistogramVec(opts, labels)

	err := prometheus.Register(vector)

	if err != nil {
		are, ok := err.(prometheus.AlreadyRegisteredError)

		if ok {
			existing := are.ExistingCollector.(prometheus.HistogramVec)
			vector = &existing
		} else {
			return nil, err
		}
	}

	return &Histo{
		vector: vector,
	}, nil
}

func (h *Histo) Observe(value float64, labels ...string) {
	h.vector.WithLabelValues(labels...).Observe(value)
}

func (h *Histo) Start() {
	h.start = time.Now()
}

func (h *Histo) Stop() {
	h.duration = time.Since(h.start)
}

func (h *Histo) Duration() float64 {
	return h.duration.Seconds()
}

func (h *Histo) Record(labels ...string) {
	h.Observe(h.Duration(), labels...)
}
